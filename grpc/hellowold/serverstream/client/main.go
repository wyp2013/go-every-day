package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "go-every-day/grpc/hellowold/serverstream/helloworld"
	"google.golang.org/grpc/metadata"
	"io"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:55556", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ctx, cancle := context.WithTimeout(context.Background(), time.Second)
	defer cancle()

	cli := pb.NewGreeterClient(conn)

	var md metadata.MD
	c, err := cli.SayHello(ctx, &pb.HelloRequest{Age:20, Name:"test"}, grpc.Header(&md))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for {
		reply, err := c.Recv()

		if err == io.EOF {
			fmt.Println("receive end")
			break
		}

		if err != nil {
			fmt.Println("receive error", err.Error())
			break
		}

		fmt.Println(reply)
	}
}