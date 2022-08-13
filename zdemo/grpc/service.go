package main

import (
	pb "Gossssss/demosss/grpc/proto"
	"context"
	"net"
	"strings"

	"google.golang.org/grpc"
)

// 必须要实现某个结构体
type Greeter struct{}

func (g *Greeter) SayHi(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	// code
	return &pb.HelloResponse{
		Msg: "hi, " + req.Name,
	}, nil
}

func (g *Greeter) GetUpperName(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Msg: strings.ToUpper(req.Name),
	}, nil
}

func main() {
	// tcp
	listener, _ := net.Listen("tcp", "127.0.0.1:8081")

	// 创建服务
	srv := grpc.NewServer()
	// Greeter注册到gRPC服务
	pb.RegisterHelloServiceServer(srv, &Greeter{})

	// 启动服务
	_ = srv.Serve(listener)
}
