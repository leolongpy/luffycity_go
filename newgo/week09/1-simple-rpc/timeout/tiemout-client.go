package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	proto2 "luffycity_go/newgo/week09/1-simple-rpc/proto"
	"time"
)

var grpcClient proto2.SimpleClient

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("err %v", err)
	}
	defer conn.Close()

	ctx := context.Background()
	grpcClient = proto2.NewSimpleClient(conn)
	route(ctx, 3)
}

func route(ctx context.Context, deadlines time.Duration) {
	// 设置3秒超时时间
	clientDeadline := time.Now().Add(time.Duration(deadlines * time.Second))
	ctx, cancel := context.WithDeadline(ctx, clientDeadline)
	defer cancel()
	// 发送结构体
	req := proto2.SimpleReq{
		Data: "grpc",
	}
	//  传入超时时间为3秒的ctx
	res, err := grpcClient.Route(ctx, &req)
	if err != nil {
		// 获取错误状态
		statu, ok := status.FromError(err)
		if ok {

			// 判断是否为调用超时
			if statu.Code() == codes.DeadlineExceeded {
				log.Fatalln("Route timeout")
			}
		}
		log.Fatalf("call route err : %v", err)
	}
	log.Println(res.Value)
}
