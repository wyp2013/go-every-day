package main

import (
	"fmt"
	"google.golang.org/grpc"
	pb "go-every-day/grpc/hellowold/cliserstream/helloworld"
	"time"
	"context"
)

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	client := pb.NewGreeterClient(conn)

	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	stream, _ := client.SayHello(ctx)

	for i := 0; i < 5; i++ {
		name := fmt.Sprintf("xx-%d-yy", i)
		age := int32(i + 10)
		req := &pb.HelloRequest{Name: name, Age: age}

		err := stream.Send(req)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}


		reply, err := stream.Recv()
		if err != nil {
			fmt.Println(reply)
		}
	}


}
