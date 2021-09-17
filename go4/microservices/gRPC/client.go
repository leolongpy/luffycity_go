package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"luffycity_go/go4/microservices/gRPC/proto"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接异常", err.Error())
	}
	defer conn.Close()
	client := proto.NewUserInfoServiceClient(conn)
	req := &proto.UserRequest{Name: "zs"}
	response, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		fmt.Println("响应异常", err.Error())
	} else {
		fmt.Printf("响应结果：%v\n", response)
	}
}
