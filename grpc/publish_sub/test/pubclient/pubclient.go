package main

import (
	"context"
	"fmt"
	pb "go-every-day/grpc/publish_sub/pubsub"
	"google.golang.org/grpc"
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

	pubClient := pb.NewPubsubServiceClient(conn)

	pubClient.Publish(ctx, &pb.PubRequest{Topic:"go learn go go go", TopicType:1})
	pubClient.Publish(ctx, &pb.PubRequest{Topic:"python python python aaa ", TopicType:2})


}
