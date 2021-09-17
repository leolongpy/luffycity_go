package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	proto2 "luffycity_go/newgo/week09/3-client-stream_rpc/proto"
	"strconv"
)

var streamClient proto2.StreamClientClient

func routeList() {
	stream, err := streamClient.RouteList(context.Background())
	if err != nil {
		log.Fatalf("Upload list err %v", err)
	}

	for n := 0; n < 5; n++ {
		err := stream.Send(&proto2.StreamReq{StreamValue: "stream clinet rpc" + strconv.Itoa(n)})
		log.Println(n)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream requset err %v", err)
		}
	}
	rsp, err := stream.CloseAndRecv()
	log.Printf("%v", rsp)
}
func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connent err :%v", err)
	}
	defer conn.Close()
	streamClient = proto2.NewStreamClientClient(conn)
	routeList()
}
