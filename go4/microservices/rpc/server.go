package rpc

import (
	"fmt"
	"log"
	"net"
	"reflect"
)

// Server 声明服务端
type Server struct {
	//  地址
	addr  string
	funcs map[string]reflect.Value
}

func NewServer(addr string) *Server {
	return &Server{
		addr:  addr,
		funcs: make(map[string]reflect.Value),
	}
}

//服务端注册
// 第一个参数函数名 第二个传入真正的函数
func (s *Server) Register(RPCName string, f interface{}) {
	if _, ok := s.funcs[RPCName]; ok {
		return
	}
	fVal := reflect.ValueOf(f)
	s.funcs[RPCName] = fVal
}

// 服务端等待调用的方法
func (s *Server) Run() {
	//监听
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("监听 %s err:%v", s.addr, err)
		return
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			return
		}
		serSession := NewSession(conn)
		b, err := serSession.Read()
		if err != nil {
			return
		}
		rpcData, err := decode(b)
		if err != nil {
			return
		}
		// 根据读到的name，得到调用的函数
		f, ok := s.funcs[rpcData.Name]
		if !ok {
			fmt.Printf("函数 %s 不存在", rpcData.Name)
			return
		}
		inArgs := make([]reflect.Value, 0, len(rpcData.Args))
		for _, arg := range rpcData.Args {
			inArgs = append(inArgs, reflect.ValueOf(arg))
		}
		// 反射调用方法
		out := f.Call(inArgs)
		outArgs := make([]interface{}, 0, len(out))
		for _, o := range out {
			outArgs = append(outArgs, o.Interface())
		}
		respData := RPCData{rpcData.Name, outArgs}
		bytes, err := encode(respData)
		if err != nil {
			return
		}
		err = serSession.Write(bytes)
		if err != nil {
			return
		}
	}
}
