package main

import (
	"fmt"
	pb "go-every-day/grpc/hellowold/clientstream/helloworld"
	"google.golang.org/grpc"
	"io"
	"net"
)

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type Server struct {
}

func (s *Server) SayHello(srv pb.Greeter_SayHelloServer) error {
	count := 0

	for {
		req, err := srv.Recv()
		count++

		if err == io.EOF {
			id := int32(count)
			return srv.SendAndClose(&pb.HelloReply{Message:"end!", Id: id,})
		}

		if err != nil {
			return err
		}

		// do something
		fmt.Println(req)
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":55551")
	defer lis.Close()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	gSer := grpc.NewServer()
	pb.RegisterGreeterServer(gSer, &Server{})

	err = gSer.Serve(lis)
	if err != nil {
		fmt.Println(err)
	}
}
