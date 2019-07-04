package main

import (
	"context"
	"fmt"
	//lb "google.golang.org/grpc/balancer"
	pb "go-every-day/grpc/hellowold/simple/helloworld"
	"google.golang.org/grpc"
	"time"
)

func main() {
	//grpc.WithBalancerName()
	//lb.Register()

	conn, err := grpc.Dial("localhost:55555", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c := pb.NewGreeterClient(conn)

	ctx, cancle := context.WithTimeout(context.Background(), time.Second)
	defer cancle()


	for i := 0; i < 5; i++ {
		name := fmt.Sprintf("xx-%d-yy", i)
		age := int32(i+10)
		req := &pb.HelloRequest{Name: name, Age: age}

		reply, err := c.SayHello(ctx, req)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Println(reply)
	}

}
