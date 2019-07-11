package main

import (
	"context"
	"crypto/tls"
	"fmt"
	gw "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go-every-day/grpc/rcp_json/certs"
	pb "go-every-day/grpc/rcp_json/proto"
	"go-every-day/grpc/rcp_json/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"net/http"
	"path"
	"runtime"
	"strings"
)


func getFileNames() (string, string) {
	_, filename, _, ok1 := runtime.Caller(0)
	if !ok1 {
		panic("No caller information")
	}
	dir := path.Dir(filename)
	permPathFile := path.Join(dir, "certs/server.pem")
	keyPathFile  := path.Join(dir, "certs/server.key")

	return permPathFile, keyPathFile
}

func NewGrpcServer(permFile, keyFile string) *grpc.Server {
	creds, err := credentials.NewServerTLSFromFile(permFile, keyFile)
	if err != nil {
		panic(err.Error())
	}

	var opts []grpc.ServerOption
	opts = append(opts, grpc.Creds(creds))
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterGreeterServer(grpcServer, &server.HelloServer{})

	return grpcServer
}

func GrpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	if otherHandler == nil {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			grpcServer.ServeHTTP(w, r)
		})
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}

func NewGatewayServer(permFile, certName string, endpoint string) http.Handler {
	// 这里其实和client的代码一样，内部rpc也要tls认证
	dcreds, err := credentials.NewClientTLSFromFile(permFile, certName)
	if err != nil {
		panic(err.Error())
	}

	ctx := context.Background()

	dopts := []grpc.DialOption{grpc.WithTransportCredentials(dcreds)}
	gwmux := gw.NewServeMux()

	/*
	RegisterGreeterHandlerFromEndpoint干的事情：
	1. 根据endpoint和认证参数，创建一个rpc client
	2. 注册gateway的hanlder到gwmux中，比如/hello_word对应的rpc接口sayHello
	 */
	if err := pb.RegisterGreeterHandlerFromEndpoint(ctx, gwmux, endpoint, dopts); err != nil {
		panic(err.Error())
		return nil
	}

	return gwmux
}

func main() {
	// 加载安全认证
	ts, err := certs.GetTLSConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 获取验证文件路径
	permPathFile, keyPathFile := getFileNames()
	fmt.Println(permPathFile, keyPathFile, ts)

	endpoint := ":8090"

	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer lis.Close()

	// grpc server
	grpcServer := NewGrpcServer(permPathFile, keyPathFile)

	// gateway server，其实就是把通过https的请求转换成rpcClient请求的一个处理handle
	gwmux := NewGatewayServer(permPathFile, "localhost", endpoint)

	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	ser := http.Server{
		Addr:     endpoint,
		Handler:   GrpcHandlerFunc(grpcServer, mux),
		TLSConfig: ts,
	}

	err = ser.Serve(tls.NewListener(lis, ts))
	if err != nil {
		panic(err.Error())
	}
}
