package rpc

import (
	"bytes"
	"encoding/gob"
)

// RPCData 定义RPC交互的数据结构
type RPCData struct {
	//访问的函数
	Name string
	// 访问时的参数
	Args []interface{}
}

// 编码
func encode(data RPCData) ([]byte, error) {
	//得到字节数组的编码器
	var buf bytes.Buffer
	bufEnc := gob.NewEncoder(&buf)
	err := bufEnc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// 解码
func decode(b []byte) (RPCData, error) {
	buf := bytes.NewBuffer(b)
	bufDec := gob.NewDecoder(buf)
	var data RPCData
	err := bufDec.Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
