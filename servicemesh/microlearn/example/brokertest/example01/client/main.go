package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
)

func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"http://127.0.0.1:2379",
			// "http://127.0.0.1:12379", "http://127.0.0.1:22379",
		}
	})

	service := micro.NewService(
		micro.Name("HelloGreeter1"),
		micro.Registry(reg),
	)

	service.Init()

	cmsg := `{"message": "Hello World"}`
	err := service.Client().Publish(context.Background(), service.Client().NewMessage("testTopic", cmsg))
	if err != nil {
		// 这个测试会报错"{"id":"go.micro.client","code":500,"detail":"Unsupported Content-Type: application/grpc+proto","status":"Internal Server Error"}
		// 原因： https://github.com/micro/go-micro/issues/625
		fmt.Println(err.Error())
	}

	msg := &broker.Message{
		Header: map[string]string{
			"Content-Type": "application/json",
		},
		Body: []byte(`{"message": "Hello World"}`),
	}

	// broker的publish是异步发送
	err = service.Options().Broker.Publish("testTopic", msg)
	if err != nil {
		fmt.Println(err.Error())
	}

	// 阻塞push信息发送完成
	select {}
}
