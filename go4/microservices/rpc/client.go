package rpc

import (
	"net"
	"reflect"
)

type Client struct {
	conn net.Conn
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		conn: conn,
	}
}

func (c *Client) callRPC(rpcName string, fPtr interface{}) {
	fn := reflect.ValueOf(fPtr).Elem()
	f := func(args []reflect.Value) []reflect.Value {
		//处理参数
		inArgs := make([]interface{}, 0, len(args))
		for _, arg := range args {
			inArgs = append(inArgs, arg.Interface())
		}
		cliSession := NewSession(c.conn)
		reqRPC := RPCData{Name: rpcName, Args: inArgs}
		b, err := encode(reqRPC)
		if err != nil {
			panic(err)
		}
		err = cliSession.Write(b)
		if err != nil {
			panic(err)
		}
		respBytes, err := cliSession.Read()
		if err != nil {
			panic(err)
		}
		//解码
		respRPC, err := decode(respBytes)
		if err != nil {
			panic(err)
		}
		//处理服务端返回的数据
		outArgs := make([]reflect.Value, 0, len(respRPC.Args))
		for i, arg := range respRPC.Args {
			//进行nil转换
			if arg == nil {
				outArgs = append(outArgs, reflect.Zero(fn.Type().Out(i)))
				continue
			}
			outArgs = append(outArgs, reflect.ValueOf(arg))
		}
		return outArgs
	}
	v := reflect.MakeFunc(fn.Type(), f)
	fn.Set(v)
}
