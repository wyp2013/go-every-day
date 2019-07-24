package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	sgrpc "github.com/micro/go-micro/server/grpc"
	cgrpc "github.com/micro/go-micro/client/grpc"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	pb "go-every-day/servicemesh/microlearn/example/hello/proto"
)

func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"http://127.0.0.1:2379",
			// "http://127.0.0.1:12379", "http://127.0.0.1:22379",
		}
	})

	service := micro.NewService(
		micro.Server(sgrpc.NewServer()),  // 这里不用默认的rpcserver，改用grpcserver，注释也可以运行
		micro.Client(cgrpc.NewClient()),  // 对应的client应该也是 grpc
		micro.Name("HelloGreeter"),
		micro.Registry(reg),
	)

	service.Init()

	ctx := context.Background()

	// 服务提供方的名字
	helloClient := pb.NewGreeterService("HelloGreeter", service.Client())

	for i := 0; i < 5; i++ {
		name := fmt.Sprintf("xx-%d-yy", i)
		age := int32(i + 10)
		req := &pb.HelloRequest{Name: name, Age: age}

		reply, err := helloClient.SayHello(ctx, req)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Println("response: ", reply)
	}

	stremClient, err := helloClient.SayHello2(ctx)
	defer stremClient.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i := 0; i < 5; i++ {
		name := fmt.Sprintf("sayhllo2-%d", i)
		age := int32(i + 10)
		req := &pb.HelloRequest{Name: name, Age: age}
		err := stremClient.Send(req)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

}
