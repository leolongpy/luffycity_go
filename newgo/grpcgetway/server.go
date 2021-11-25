package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	proto2 "simple/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)
var port = flag.Int("port", 50051, "the port to serve on")
var restful = flag.Int("restful", 8080, "the port to restful serve on")

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
	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	proto2.RegisterSimpleServer(s, &SimpleServer{})
	// Serve gRPC Server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Serving gRPC on 0.0.0.0" + fmt.Sprintf(":%d", *port))
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()




	gwmux := runtime.NewServeMux()
	proto2.RegisterSimpleHandlerFromEndpoint(context.Background(),gwmux,"localhost:50051",[]grpc.DialOption{grpc.WithInsecure()})
	http.HandleFunc("/healthz", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("ok"))
	})
	http.Handle("/", gwmux)

	log.Println("Serving gRPC-Gateway on http://0.0.0.0"+fmt.Sprintf(":%d", *restful))
	log.Fatalln(http.ListenAndServe(":8080",nil))
}
