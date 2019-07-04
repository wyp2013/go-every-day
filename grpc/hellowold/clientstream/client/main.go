package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "go-every-day/grpc/hellowold/clientstream/helloworld"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:55551", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	client := pb.NewGreeterClient(conn)
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	stream, err := client.SayHello(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i := 0; i < 5; i++ {
		name := fmt.Sprintf("xx-%d-yy", i)
		age := int32(i + 10)
		req := &pb.HelloRequest{Name: name, Age: age}

		err := stream.Send(req)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}


	}

	replay, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(replay)
}
