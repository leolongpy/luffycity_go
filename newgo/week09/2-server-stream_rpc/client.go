package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	proto2 "luffycity_go/newgo/week09/2-server-stream_rpc/proto"
)

func main() {
	req := proto2.SimpleReq{Data: "stream server grpc"}
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("net connect err :%v", err)
	}
	defer conn.Close()
	grpcClient := proto2.NewStreamServerClient(conn)
	stream, err := grpcClient.ListValue(context.Background(), &req)
	if err != nil {
		log.Fatalf("call ListValue err :%v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("get stream err :%v", err)
		}
		log.Println(res.StreamValue)
		break
	}
	//stream.CloseSend()
	//res, _ := stream.Recv()
	//log.Println(res.StreamValue)
	//res1, _ := stream.Recv()
	//log.Println(res1.StreamValue)
}
