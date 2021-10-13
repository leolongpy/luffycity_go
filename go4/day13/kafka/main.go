package main

import (
	"fmt"
	//"gopkg.in/Shopify/sarama.v1"
	"github.com/Shopify/sarama"
)
func main() {
	//1.生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll//ACK
	config.Producer.Partitioner = sarama.NewRandomPartitioner//分区
	config.Producer.Return.Successes = true//确认
	//2.链接kafka
	client,err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"},config)
	if err != nil {
		fmt.Println("producer closed,err:",err)
		return
	}
	defer client.Close()

	//3.封装消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "today"
	msg.Value = sarama.StringEncoder("2021.09.20")

	//4.发送消息
	pid,offset,err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg faild,err:",err)
	}
	fmt.Printf("pid:%v offset:%v\n",pid,offset)
}
//.\kafka-console-consumer.bat --bootstrap-server 127.0.0.1:9092 --topic today --from-beginning
// .\kafka-server-start.bat ..\..\config\server.properties
// .\zookeeper-server-start.bat ..\..\config\zookeeper.properties