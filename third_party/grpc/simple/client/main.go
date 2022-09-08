package client

import (
	pb "Gossssss/third_party/grpc/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	// 服务地址
	conn, _ := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	defer conn.Close()

	// 获取gRPC客户端
	client := pb.NewHelloServiceClient(conn)

	// 远程过程调用
	response1, _ := client.SayHi(

		context.Background(),
		&pb.HelloRequest{Name: "yao"})

	fmt.Println(response1.Msg)

	response2, _ := client.GetUpperName(
		context.Background(), &pb.HelloRequest{
			Name: "yao",
		})
	fmt.Println(response2.Msg)
}
