package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	pb "go-every-day/servicemesh/microlearn/example/hello/proto"
)

func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"http://127.0.0.1:2379", "http://127.0.0.1:12379", "http://127.0.0.1:22379",
		}
	})

	service := micro.NewService(
		micro.Name("HelloGreeter1"),
		micro.Registry(reg),
	)

	service.Init()

	// 服务提供方的名字
	helloClient := pb.NewGreeterService("HelloGreeter", service.Client())

	for i := 0; i < 5; i++ {
		name := fmt.Sprintf("xx-%d-yy", i)
		age := int32(i + 10)
		req := &pb.HelloRequest{Name: name, Age: age}

		reply, err := helloClient.SayHello(context.Background(), req)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Println("response: ", reply)
	}
}
