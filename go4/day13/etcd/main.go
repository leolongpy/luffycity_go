package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

// 代码连接etcd

func main(){
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
		DialTimeout:time.Second*5,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v", err)
		return
	}

	defer cli.Close()

	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	str := `[{"path":E:/desktop/log1.log","topic":"log1"},{"path":"E:/desktop/log2.log","topic":"log2"}]`
	_, err = cli.Put(ctx, "log_192.168.0.106", str)
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v", err)
		return
	}
	cancel()

	// get
	//ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	//gr, err := cli.Get(ctx, "collect_log_conf")
	//if err != nil {
	//	fmt.Printf("get from etcd failed, err:%v", err)
	//	return
	//}
	//for _, ev := range gr.Kvs{
	//	fmt.Printf("key:%s value:%s\n", ev.Key, ev.Value)
	//}
	//cancel()
}
