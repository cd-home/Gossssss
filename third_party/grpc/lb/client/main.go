package main

import (
	pb "Gossssss/third_party/grpc/api"
	"Gossssss/third_party/grpc/lb/etcdv3/discovery"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
	"log"
	"strings"
)

var endpoints string
var service string

func init() {
	flag.StringVar(&endpoints, "endpoints", "127.0.0.1:2379", "endpoints")
	flag.StringVar(&service, "service", "hello", "service")
	flag.Parse()
}

func main() {
	_endpoints := strings.Split(endpoints, ",")
	r := discovery.New(_endpoints)
	resolver.Register(r)
	conn, err := grpc.Dial(
		// URI schemes
		// dns:[//authority/]host[:port] -- DNS (default)
		r.Scheme()+"://authority/"+service,
		// grpc.WithBalancerName() 废弃
		grpc.WithDefaultServiceConfig(
			fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`,
				roundrobin.Name)),
		grpc.WithInsecure())
	if err != nil {
		return
	}
	grpcClient := pb.NewHelloServiceClient(conn)
	response, err := grpcClient.SayHi(context.Background(), &pb.HelloRequest{
		Name: "yao",
	})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(response.Msg)
}
