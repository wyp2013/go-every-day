package rpc

import (
	"context"
	"fmt"
	pb "go-every-day/grpc/publish_sub/pubsub"
	"time"
)

type PubsubServer struct {
	publish Publish
}

func NewPubsubServer(timeout time.Duration, bufferSize int) *PubsubServer{
	return &PubsubServer{
		publish: NewPublisher(timeout, bufferSize),
	}
}

func (m *PubsubServer) Publish(ctx context.Context, req *pb.PubRequest) (*pb.PubReply, error) {
	fmt.Println("", req.TopicType, req.Topic)

	m.publish.PublishMessage(req.Topic)

	return &pb.PubReply{Content: "success", Errno: 0}, nil
}

func (m *PubsubServer) SubscribeTopic(req *pb.SubRequest, srv pb.PubsubService_SubscribeTopicServer) error {
	var chanMsg chan interface{}

	// 订阅 go
	if req.TopicType == 1 {
		chanMsg = m.publish.SubscribeTopic(subGo)
	} else if req.TopicType == 2 {
		chanMsg = m.publish.SubscribeTopic(subPython)
	} else {
		chanMsg = m.publish.SubscribeTopic(nil)
	}

	// go func() {
		for msg := range chanMsg {
			fmt.Println("chan ", msg)

			reply := &pb.SubReply{Content: msg.(string), Errno:0}
			err := srv.Send(reply)

			if err != nil {
				fmt.Println("send error: ", err.Error())
				pub := (m.publish).(*Publisher)
				pub.Retire(chanMsg)
				break
			}
		}
	// }()   // 用 gorouting 就会有错误

	return nil
}


