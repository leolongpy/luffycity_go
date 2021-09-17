package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	proto2 "luffycity_go/newgo/week09/1-simple-rpc/proto"
	"net"
)

type SimpleServer struct {
	proto2.UnimplementedSimpleServer
}

func (s *SimpleServer) Route(ctx context.Context, req *proto2.SimpleReq) (*proto2.SimpleRsp, error) {
	res := proto2.SimpleRsp{
		Code:  200,
		Value: "hello" + req.Data,
	}
	return &res, nil
}

func main() {
	listenner, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("net listen err :%v", err)
	}
	log.Println("8080 net.Listing")
	grpcServer := grpc.NewServer()
	proto2.RegisterSimpleServer(grpcServer, &SimpleServer{})
	err = grpcServer.Serve(listenner)
	if err != nil {
		log.Fatalf("grpcServer.Server err:%v", err)
	}
}
