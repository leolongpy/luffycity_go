package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Server struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIp"`
}

type Servers struct {
	ServerName string   `xml:"serverName"`
	Version    int      `xml:"version"`
	Servers    []Server `xml:"servers"`
}

func main() {
	data, err := ioutil.ReadFile("./my.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	var servers Servers
	fmt.Println(data)
	err = xml.Unmarshal(data, &servers)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("xml:%#v\n", servers)
}
