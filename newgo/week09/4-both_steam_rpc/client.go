package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	proto2 "luffycity_go/newgo/week09/4-both_steam_rpc/proto"
	"strconv"
)

var streamClient proto2.StreamClient

func conversations() {
	stream, err := streamClient.Conversations(context.Background())
	if err != nil {
		log.Fatalf("err : %v", err)
	}

	for n := 0; n < 5; n++ {
		err := stream.Send(&proto2.StreamReq{
			Question: "stream client rpc " + strconv.Itoa(n),
		})
		if err != nil {
			log.Fatalf("stream request err: %v", err)
		}
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Conversation get stream err :%v", err)
		}
		log.Println(res.Answer)
	}
	err = stream.CloseSend()
	if err != nil {
		log.Fatalf("Conversation close stream err: %v", err)
	}
}
func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	streamClient = proto2.NewStreamClient(conn)
	conversations()
}
