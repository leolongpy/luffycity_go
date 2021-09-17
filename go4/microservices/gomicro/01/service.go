package main

import (
	"fmt"
	"github.com/asim/go-micro/v3"
	"golang.org/x/net/context"
	"log"
	"luffycity_go/go4/microservices/gomicro/01/proto"
)

type Hello struct {
}

func (h *Hello) Info(ctx context.Context, req *proto.InfoRequest, resp *proto.InfoResponse) error {
	resp.Msg = "你好" + req.Username
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("hello"),
	)

	service.Init()
	err := proto.RegisterHelloHandler(service.Server(), new(Hello))
	if err != nil {
		fmt.Println(err.Error())
	}
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
