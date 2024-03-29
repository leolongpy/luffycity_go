package etcd

import (
	"context"
	"fmt"
	"luffycity_go/go6/registry"

	"testing"
	"time"
)

func TestRegister(t *testing.T) {
	registryInst, err := registry.InitRegistry(context.TODO(), "etcd",
		registry.WithAddrs([]string{"127.0.0.1:2379"}),
		registry.WithTimeout2(time.Second),
		registry.WithRegistryPath("/my/etcd/"),
		registry.WithHeartBeat(5),
	)
	if err != nil {
		t.Errorf("init registry failed, err:%v", err)
		return
	}
	service := &registry.Service{
		Name: "comment_service",
	}
	service.Nodes = append(service.Nodes, &registry.Node{
		IP:   "127.0.0.1",
		Port: 8801,
	},
		&registry.Node{
			IP:   "127.0.0.2",
			Port: 8801,
		},
	)
	registryInst.Registry(context.TODO(), service)
	go func() {
		time.Sleep(time.Second * 10)
		service.Nodes = append(service.Nodes, &registry.Node{
			IP:   "127.0.0.3",
			Port: 8801,
		},
			&registry.Node{
				IP:   "127.0.0.4",
				Port: 8801,
			},
		)
		registryInst.Registry(context.TODO(), service)
	}()
	for {
		service, err := registryInst.GetService(context.TODO(), "comment_service")
		if err != nil {
			t.Errorf("get service failed, err:%v", err)
			return
		}

		for _, node := range service.Nodes {
			fmt.Printf("service:%s, node:%#v\n", service.Name, node)
		}
		fmt.Println()
		time.Sleep(time.Second * 5)
	}
}
