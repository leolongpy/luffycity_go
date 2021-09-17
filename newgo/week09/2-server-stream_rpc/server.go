package main

import (
	"google.golang.org/grpc"
	"log"
	proto2 "luffycity_go/newgo/week09/2-server-stream_rpc/proto"
	"net"
	"strconv"
	"time"
)

type StreamServer struct {
	proto2.UnimplementedStreamServerServer
}

func (s *StreamServer) ListValue(req *proto2.SimpleReq, srv proto2.StreamServer_ListValueServer) error {
	for n := 0; n < 15; n++ {
		err := srv.Send(&proto2.StreamRsp{StreamValue: req.Data + strconv.Itoa(n)})
		if err != nil {
			return err
		}
		log.Println(n)
		time.Sleep(1 * time.Second)
	}
	return nil
}

func main() {
	listenner, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("net listen err :%v", err)
	}
	log.Println("8080 net.Listing")
	grpcServer := grpc.NewServer()
	proto2.RegisterStreamServerServer(grpcServer, &StreamServer{})
	err = grpcServer.Serve(listenner)
	if err != nil {
		log.Fatalf("grpcServer.Server err:%v", err)
	}
}
