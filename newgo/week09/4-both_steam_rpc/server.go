package main

import (
	"google.golang.org/grpc"
	"io"
	"log"
	proto2 "luffycity_go/newgo/week09/4-both_steam_rpc/proto"
	"net"
	"strconv"
)

type StreamServer struct {
	proto2.UnimplementedStreamServer
}

func (s *StreamServer) Conversations(srv proto2.Stream_ConversationsServer) error {
	n := 1
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		err = srv.Send(&proto2.StreamRsp{
			Answer: "from steam server answeer: the " + strconv.Itoa(n) + " question is " + req.Question,
		})
		if err != nil {
			return err
		}
		n++
		log.Printf("from stream client question: %s", req.Question)
	}
}

func main() {
	listenner, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("net listen err :%v", err)
	}
	log.Println("8080 net.Listing")
	grpcServer := grpc.NewServer()
	proto2.RegisterStreamServer(grpcServer, &StreamServer{})
	err = grpcServer.Serve(listenner)
	if err != nil {
		log.Fatalf("grpcServer.Server err:%v", err)
	}
}
