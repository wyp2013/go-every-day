package main

import (
	"errors"
	"fmt"
	pb "go-every-day/grpc/hellowold/serverstream/helloworld"
	"google.golang.org/grpc"
	"math/rand"
	"net"
)

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type MyGreeterServer struct {
}

func (s *MyGreeterServer) SayHello(req *pb.HelloRequest, srv pb.Greeter_SayHelloServer) error {
	if req == nil {
		return errors.New("请求参数为空")
	}

	fmt.Println(req)

	for i := req.Age; i < req.Age+10; i++ {
		err := srv.Send(&pb.HelloReply{
			Message: fmt.Sprintf("test-rangd-%d-age-%d", rand.Int31(), i),
			Id: rand.Int31(),
			Age: rand.Int31n(27),
		})

		if err != nil {
			fmt.Println(err)
		}
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":55556")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer lis.Close()

	gSer := grpc.NewServer()
	pb.RegisterGreeterServer(gSer, &MyGreeterServer{})

	err = gSer.Serve(lis)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
