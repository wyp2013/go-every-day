package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/client"
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

	myClient := cgrpc.NewClient(client.Registry(reg))
	myClient.Options().Broker.Init(broker.Registry(reg))

	// 服务提供方的名字
	helloClient := pb.NewGreeterService("HelloGreeter", myClient)
	ctx := context.Background()

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
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer stremClient.Close()

	for i := 0; i < 5; i++ {
		name := fmt.Sprintf("sayhllo2-%d", i)
		age := int32(i + 10)
		req := &pb.HelloRequest{Name: name, Age: age}
		err := stremClient.Send(req)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	msg := &broker.Message{
		Header: map[string]string{
			"Content-Type": "application/json",
		},
		Body: []byte(`{"message": "Hello World"}`),
	}

	// broker的publish是异步发送
	err = myClient.Options().Broker.Publish("testTopic", msg)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("publish message")
	}

	// 阻塞push信息发送完成
	select {}
}
