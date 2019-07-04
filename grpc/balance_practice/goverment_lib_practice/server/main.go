package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	pb "go-every-day/grpc/balance_practice/goverment_lib_practice/simplecache"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	hv1 "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"
	"math/rand"
	"net"
	"os"
)

// UnimplementedSimpleCacheServer can be embedded to have forward compatible implementations.
type MySimpleCacheServer struct {
	hostname string
	cache map[int]string
}


func NewMySimpleCacheServer() *MySimpleCacheServer {
	return &MySimpleCacheServer{
		cache: make(map[int]string,),
	}
}

func (s *MySimpleCacheServer) Init() {
	test := "my-test-rand-"
	for i := 0; i < 100; i++ {
		s.cache[i] = fmt.Sprintf("%s%d", test, rand.Int())
	}
}

func (s *MySimpleCacheServer) set(id int, value string) bool {
	s.cache[id] = value

	return false
}

func (s *MySimpleCacheServer) GetValue(ctx context.Context, req *pb.CacheRequest) (*pb.CacheReply, error) {
	md := metadata.Pairs("server", s.hostname) // 暴露ip
	grpc.SendHeader(ctx, md)

	fmt.Println(req)

	if req == nil {
		return nil, errors.New("request is empty")
	}

	replay := new (pb.CacheReply)
	value, ok := s.cache[int(req.Id)]
	if !ok {
		if req.Set {
			s.set(int(req.Id), req.DefaultValue)
			replay.Set = 2
			replay.Value = req.DefaultValue
		} else {
			replay.Set = 0
		}
	} else {
		replay.Value = value
		replay.Set = 1
	}

	return replay, nil
}



func main() {
	port := flag.String("port", "55550", "http listen port")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", *port))
	if err != nil {
		fmt.Println(err.Error())
	}
	defer lis.Close()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	host, err := os.Hostname()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	myCache := NewMySimpleCacheServer()
	myCache.Init()
	myCache.hostname = fmt.Sprintf("%s:%s", host, *port)

	gSer := grpc.NewServer()
	pb.RegisterSimpleCacheServer(gSer, myCache)

	healthServer := health.NewServer()
	hv1.RegisterHealthServer(gSer, healthServer)
	healthServer.SetServingStatus("simplecache.SimpleCache", hv1.HealthCheckResponse_SERVING)

	err = gSer.Serve(lis)
	if err != nil {
		fmt.Println(err)
	}

}
