package main

import (
	"context"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/errors"
	_ "github.com/micro/go-plugins/registry/consul"
	"log"
	"luffycity_go/go4/microservices/gomicro/02/proto"
)

type Example struct {
}
type Foo struct {
}

func (e *Example) Call(ctx context.Context, req *proto.CallRequest, resp *proto.CallResponse) error {
	log.Println("收到Example.Call请求")
	if len(req.Name) == 0 {
		return errors.BadRequest("example", "no name")
	}
	resp.Message = "Exmaple.Call接收到了你的请求" + req.Name
	return nil
}

func (f *Foo) Bar(ctx context.Context, req *proto.EmptyRequest, resp *proto.EmptyResponse) error {
	log.Println("收到Foo.Bar请求")
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("example"),
	)

	service.Init()
	err := proto.RegisterExampleHandler(service.Server(), new(Example))
	if err != nil {
		println(err.Error())
	}
	err = proto.RegisterFooHandler(service.Server(), new(Foo))
	if err != nil {
		println(err.Error())
	}
	err = service.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}
