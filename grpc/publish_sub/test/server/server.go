package main

import (
	"fmt"
	pb "go-every-day/grpc/publish_sub/pubsub"
	"go-every-day/grpc/publish_sub/rpc"
	"google.golang.org/grpc"
	"net"
	"time"
)


func main() {
	lis, err := net.Listen("tcp", "localhost:9991")
	defer lis.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	pubsubSer := rpc.NewPubsubServer(10*time.Second, 10)

	grpcServer := grpc.NewServer()
	pb.RegisterPubsubServiceServer(grpcServer, pubsubSer)

	err = grpcServer.Serve(lis)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
