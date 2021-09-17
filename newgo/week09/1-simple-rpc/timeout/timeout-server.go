package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	proto2 "luffycity_go/newgo/week09/1-simple-rpc/proto"
	"net"
	"time"
)

type SimpleServerive struct {
	proto2.UnimplementedSimpleServer
}

func main() {

	listenner, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Println("8080 net.Listing....")

	grpcServer := grpc.NewServer()
	proto2.RegisterSimpleServer(grpcServer, &SimpleServerive{})

	err = grpcServer.Serve(listenner)
	if err != nil {
		log.Fatalf("%v", err)
	}

}

func (s *SimpleServerive) Route(ctx context.Context, req *proto2.SimpleReq) (*proto2.SimpleRsp, error) {
	select {
	case <-ctx.Done():
		log.Println(ctx.Err())
		return nil, status.Errorf(codes.Canceled, "Client cancelled, ")
	case <-time.After(4 * time.Second): // 模拟耗时操作
		res := proto2.SimpleRsp{
			Code:  200,
			Value: "hello " + req.Data,
		}
		return &res, nil
	}
}
