package main

import (
	"context"
	"fmt"
	pb "go-every-day/grpc/rcp_json/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"path"
	"runtime"
	"strings"
	"time"
)

func main() {

	_, filename, _, ok1 := runtime.Caller(0)
	if !ok1 {
		panic("No caller information")
	}
	dir := path.Dir(filename)

	dir = dir[:strings.LastIndex(dir, "/")]
	permPathFile := path.Join(dir, "certs/server.pem")
	fmt.Println(permPathFile)

	creds, err := credentials.NewClientTLSFromFile(permPathFile, "localhost")
	if err != nil {
		log.Println("Failed to create TLS credentials %v", err)
	}
	conn, err := grpc.Dial(":8090", grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Println("err ", err.Error())
		return
	}

	c := pb.NewGreeterClient(conn)

	ctx, cancle := context.WithTimeout(context.Background(), time.Second)
	defer cancle()


	for i := 0; i < 5; i++ {
		name := fmt.Sprintf("xx-%d-yy", i)
		age := int32(i+10)
		req := &pb.HelloRequest{Name: name, Age: age}

		reply, err := c.SayHello(ctx, req)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Println(reply)
	}

}
