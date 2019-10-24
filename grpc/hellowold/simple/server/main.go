package main

import (
	"context"
	"errors"
	"fmt"
	pb "go-every-day/grpc/hellowold/simple/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"math/rand"
	"net"
)

type HelloServer struct {

}

func (s *HelloServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	if in == nil {
		return nil, errors.New("input is empty")
	}

	age := in.GetAge()
	name := in.GetName()
	id := rand.Int31()

	// test send header
	md := metadata.MD{}
	md["test"] = []string{"hello"}
	grpc.SendHeader(ctx, md)

	return &pb.HelloReply{
		Message: fmt.Sprintf("you name is %s, age is %d, give you id is %d", name, age, id),
		Id: id,
		Age: age,
	}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":55555")
	gSer := grpc.NewServer(grpc.UnknownServiceHandler(nil))
	pb.RegisterGreeterServer(gSer, &HelloServer{})

	err := gSer.Serve(lis)
	if err != nil {
		fmt.Println(err.Error())
	}
}
