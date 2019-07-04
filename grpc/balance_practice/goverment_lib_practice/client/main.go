package main

import (
	"context"
	"flag"
	"fmt"
	rs "go-every-day/grpc/balance_practice/goverment_lib_practice/client/reslover"
	pb "go-every-day/grpc/balance_practice/goverment_lib_practice/simplecache"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"time"
)

const (
	defaultServer       = ":9001"
	defaultTimeout      = time.Second * 20
	defaultWait         = time.Second * 3
	defaultN            = 10
	defaultServerAddr   = "127.0.0.1:9001,127.0.0.1:9002,127.0.0.1:9003"
	defaultResolverType = "manual"
)


func healthWatch(ctx context.Context, service string, client health.HealthClient) {
	c, err := client.Watch(ctx, &health.HealthCheckRequest{Service:"simplecache.SimpleCache"})
	if err != nil {
		fmt.Println("wathc error ", err.Error())
		return
	}

	for {
		rep, err := c.Recv()
		//fmt.Println("watch ", c.Header())

		if err == io.EOF {
			fmt.Println("over")
			return
		}

		if err != nil {
			fmt.Println("wathc error ", err.Error())
			break
		} else {
			fmt.Println("watch rep ", rep.String())
		}
	}
}

func main() {
	server    := flag.String("server", defaultServer, "Name or IP of the target server, including port number.")
	enableLB  := flag.Bool("enable-balance", false, "Set to true to enable client-side load balancing")
	serverIPs := flag.String("server-ipv4", defaultServerAddr, "If load balancing is enabled, this is a list of comma-separated server addresses used by the GRPC name resolver")
	resolver  := flag.String("resolver", defaultResolverType, "The resolver to use. Supported values: dns manual")
	flag.Parse()

	opts := []grpc.DialOption{grpc.WithInsecure()}

	if *enableLB {
		opts = append(opts, grpc.WithBalancerName(roundrobin.Name))

		rt, err := rs.ParseResolverType(*resolver)
		if err != nil {
			fmt.Println(err.Error())
		}

		rs.RegisterResolver(rt, *serverIPs)
	}

	fmt.Println(*enableLB)

	conn, err := grpc.Dial(fmt.Sprintf("%s", *server), opts...)
	if err != nil {
		log.Fatal(err)
		fmt.Println("yyyyyyyyyyyyyyyyyyyyyyyyyyyy ", err.Error())
		return
	}
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := pb.NewSimpleCacheClient(conn)

	healthClient := health.NewHealthClient(conn)
	go healthWatch(ctx, "simplecache.SimpleCache", healthClient)

	for i := 0; i < 1000; i++ {

		var md metadata.MD
		reply, err := c.GetValue(ctx, &pb.CacheRequest{Id: int32(i), DefaultValue: fmt.Sprintf("test-%d", i), Set:true}, grpc.Header(&md))
		if err != nil {
			fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx ", err.Error())
		}

		fmt.Println(md, reply)

		time.Sleep(100*time.Millisecond)
	}

	fmt.Println("over")
}
