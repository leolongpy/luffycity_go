package main

import (
	"fmt"
	"log"
)

type DataSource interface {
	Push(data string)
	Query(name string) string
}
type redis struct {
	Name string
	Addr string
}

func (r *redis) Push(data string) {
	log.Printf("[Pushdata][ds.name:%s][data:%s]", r.Name, data)
}
func (r *redis) Query(name string) string {
	log.Printf("[Query.data][ds.name:%s][name:%s]", r.Name, name)
	return name + "_" + r.Name
}

type kafka struct {
	Name string
	Addr string
}

func (k *kafka) Push(data string) {
	log.Printf("[Pushdata][ds.name:%s][data:%s]", k.Name, data)
}
func (k *kafka) Query(name string) string {
	log.Printf("[Query.data][ds.name:%s][name:%s]", k.Name, name)
	return name + "_" + k.Name
}

var DataSourceManager = make(map[string]DataSource)

func register(name string, ds DataSource) {
	DataSourceManager[name] = ds
}
func main() {
	r := redis{
		Name: "redis-6.0",
		Addr: "1.1",
	}
	k := kafka{
		Name: "kafka-2.11",
		Addr: "2.2",
	}
	register("redis", &r)
	register("kafka", &k)
	//模拟推送数据
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key_%d", i)
		for _, ds := range DataSourceManager {
			ds.Push(key)
		}
	}
	//模拟查询数据
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key_%d", i)
		for _, ds := range DataSourceManager {
			log.Println(ds.Query(key))
		}
	}
}
