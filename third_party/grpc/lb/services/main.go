package main

import (
	pb "Gossssss/third_party/grpc/api"
	"Gossssss/third_party/grpc/lb/etcdv3/register"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
)

var endpoints string
var service string
var lease int64
var host string
var port string

func init() {
	flag.StringVar(&endpoints, "endpoints", "127.0.0.1:2379", "endpoints")
	flag.StringVar(&host, "host", "127.0.0.1", "host")
	flag.StringVar(&port, "port", "8080", "port")
	flag.StringVar(&service, "service", "hello", "service")
	flag.Int64Var(&lease, "lease", 5, "lease")
	flag.Parse()
}

type Greeter struct{}

func (g *Greeter) SayHi(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	// code
	return &pb.HelloResponse{
		Msg: "hi, " + req.Name + ", From: " + port,
	}, nil
}

func (g *Greeter) GetUpperName(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Msg: strings.ToUpper(req.Name),
	}, nil
}

func main() {
	_endpoints := strings.Split(endpoints, ",")
	addr := fmt.Sprintf("%s:%s", host, port)
	sr, err := register.New(_endpoints, service, addr, lease)
	if err != nil {
		log.Fatal(err)
	}
	defer sr.Revoke()

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &Greeter{})
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
