package main

import (
	"context"
	"fmt"
	pb "go-every-day/grpc/publish_sub/pubsub"
	"google.golang.org/grpc"
	"io"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:9991", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ctx, cancle := context.WithTimeout(context.Background(), 60 * time.Second)
	defer cancle()

	subClient := pb.NewPubsubServiceClient(conn)

	stream, err := subClient.SubscribeTopic(ctx, &pb.SubRequest{TopicType: 1, Topic: "go"})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Println("end")
				return
			}

			fmt.Println(err.Error())
		}


		fmt.Println("receive: ", reply.Content)
	}
}
