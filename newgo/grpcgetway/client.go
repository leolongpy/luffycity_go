package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	proto2 "simple/proto"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err : %v", err)
	}
	defer conn.Close()

	grpcClient := proto2.NewSimpleClient(conn)
	req := proto2.SimpleReq{Data: "grpc"}
	res, err := grpcClient.Route(context.Background(), &req)
	if err != nil {
		log.Fatalf("call route err : %v", err)
	}
	log.Println(res)
}
