package main

import (
	"google.golang.org/grpc"
	"io"
	"log"
	proto2 "luffycity_go/newgo/week09/3-client-stream_rpc/proto"
	"net"
)

type StreamServer struct {
	proto2.UnimplementedStreamClientServer
}

func (s *StreamServer) RouteList(srv proto2.StreamClient_RouteListServer) error {
	for {
		res, err := srv.Recv()
		if err == io.EOF {
			return srv.SendAndClose(&proto2.SimpleRsp{Value: "ok"})
		}
		if err != nil {
			return err
		}
		log.Println(res.StreamValue)
	}
}

func main() {
	listenner, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("net listen err :%v", err)
	}
	log.Println("8080 net.Listing")
	grpcServer := grpc.NewServer()
	proto2.RegisterStreamClientServer(grpcServer, &StreamServer{})
	err = grpcServer.Serve(listenner)
	if err != nil {
		log.Fatalf("grpcServer.Server err:%v", err)
	}
}
