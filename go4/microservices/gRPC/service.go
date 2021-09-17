package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"luffycity_go/go4/microservices/gRPC/proto"
	"net"
)

//定义空接口
type UserInfoService struct {
	proto.UnimplementedUserInfoServiceServer
}

// 实现方法
func (u *UserInfoService) GetUserInfo(ctx context.Context, req *proto.UserRequest) (resp *proto.UserResponse, err error) {
	name := req.Name
	if name == "zs" {
		resp = &proto.UserResponse{
			Id:    1,
			Name:  name,
			Age:   22,
			Hobby: []string{"sing", "Run"},
		}
	}
	return
}

func main() {
	addr := ":8080"
	linstener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("监听异常：%s\n", err.Error())
	}
	fmt.Printf("监听端口：%s\n", addr)
	s := grpc.NewServer()
	proto.RegisterUserInfoServiceServer(s, &UserInfoService{})
	s.Serve(linstener)

}

//protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative user.proto
