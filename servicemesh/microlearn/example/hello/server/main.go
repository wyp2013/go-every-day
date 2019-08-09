package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server/grpc"
	"github.com/micro/go-plugins/registry/etcdv3"
	pb "go-every-day/servicemesh/microlearn/example/hello/proto"
	"io"
	"math/rand"
)

type GreeterX struct {
}

// Sends a greeting
func (s *GreeterX) SayHello(ctx context.Context, in *pb.HelloRequest, reply *pb.HelloReply) error {
	if in == nil {
		return errors.New("input is empty")
	}

	age := in.GetAge()
	name := in.GetName()
	id := rand.Int31()

	reply.Message = fmt.Sprintf("you name is %s, age is %d, give you id is %d", name, age, id)
	reply.Id = id
	reply.Age = age

	return nil
}

func (s *GreeterX) SayHello2(ctx context.Context, stream pb.Greeter_SayHello2Stream) error {
	count := 0

	for {
		req, err := stream.Recv()
		count++

		if err == io.EOF {
			fmt.Println("strem is close...")
			return stream.Close()
		}

		if err != nil {
			return err
		}

		// do something
		fmt.Println("req ", req)
	}

	return nil
}


func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options){
		op.Addrs = []string{
			"http://127.0.0.1:2379",
			//"http://127.0.0.1:12379",
			//"http://127.0.0.1:22379",
		}
	})

	service := micro.NewService(
		micro.Server(grpc.NewServer()), // 这里不用默认的rpcserver，改用grpcserver，注释也可以运行
		micro.Name("HelloGreeter"),
		micro.Version("latest"),
		micro.Registry(reg),
	)

	service.Init()
	pb.RegisterGreeterHandler(service.Server(), new(GreeterX))


    // 启动broker服务
	service.Options().Broker.Connect()
	// 测试broker, Broker.Subscribe必须运行在 Broker.Connect()之后
	_, err0 := service.Options().Broker.Subscribe("testTopic", func(p broker.Event) error {
		fmt.Println("sub0 receive: ", string(p.Message().Body))
		return nil
	})
	if err0 != nil {
		fmt.Println("sub0 error: ", err0.Error())
	}

	// 启动服务
	if err := service.Run(); err != nil {
		panic(err)
	}
}
