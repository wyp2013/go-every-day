package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"http://127.0.0.1:2379",
			//"http://127.0.0.1:12379",
			//"http://127.0.0.1:22379",
		}
	})

	service := micro.NewService(
		micro.Name("testbroker"),
		micro.Version("latest"),
		micro.Registry(reg),
	)

	service.Init()

	//// 测试订阅发布模式
	// gServer := service.Server()
	//sub0 := gServer.NewSubscriber("testTopic", func(ctx context.Context, p broker.Event) error {
	//	fmt.Println("sub0 receive: ", string(p.Message().Body))
	//	return nil
	//})
	//err := gServer.Subscribe(sub0)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//
	//// 测试broker
	//sub1 := gServer.NewSubscriber("testTopic", func(ctx context.Context, p broker.Event) error {
	//	fmt.Println("sub1 receive: ", string(p.Message().Body))
	//	return nil
	//})
	//gServer.Subscribe(sub1)

	//if err := service.Run(); err != nil {
	//	panic(err)
	//}

	// 启动broker服务
	err := service.Options().Broker.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer service.Options().Broker.Disconnect()

	// 测试broker, Broker.Subscribe必须运行在 Broker.Connect()之后
	sub0, err0 := service.Options().Broker.Subscribe("testTopic", func(p broker.Event) error {
		fmt.Println("sub0 receive: ", string(p.Message().Body))
		return nil
	})
	if err0 != nil {
		fmt.Println("sub0 error: ", err0.Error())
	}
	defer sub0.Unsubscribe() // 删除注册broker


	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	select {
	case <-ch:
		break
	}
}
