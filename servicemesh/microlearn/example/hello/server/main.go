package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	pb "go-every-day/servicemesh/microlearn/example/hello/proto"
	"math/rand"
)

type Greeter struct {
}

// Sends a greeting
func (s *Greeter) SayHello(ctx context.Context, in *pb.HelloRequest, reply *pb.HelloReply) error {
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


func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options){
		op.Addrs = []string{
			"http://127.0.0.1:2379", "http://127.0.0.1:12379", "http://127.0.0.1:22379",
		}
	})

	service := micro.NewService(
		micro.Name("HelloGreeter"),
		micro.Version("latest"),
		micro.Registry(reg),
	)

	service.Init()
	pb.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		panic(err)
	}
}
