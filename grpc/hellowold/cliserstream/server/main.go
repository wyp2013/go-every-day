package main

import (
	pb "go-every-day/grpc/hellowold/cliserstream/helloworld"
	"google.golang.org/grpc"
	"fmt"
	"net"
)

type MyGreeterServer struct {
}

func (m *MyGreeterServer) SayHello(srv pb.Greeter_SayHelloServer) error {
	req, err := srv.Recv()
	if err != nil {
		return err
	}

	reply := &pb.HelloReply{
		Message: req.Name + " xxxxx",
		Id: req.Age,
		Age: req.Age,
	}

	err = srv.Send(reply)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer lis.Close()

	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &MyGreeterServer{})

	err = grpcServer.Serve(lis)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
