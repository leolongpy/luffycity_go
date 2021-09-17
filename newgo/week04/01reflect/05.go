package main

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"reflect"
)

type Person5 struct {
	Name string     `json:"name" yaml:"yaml_name"  xiaoyi:"name" `
	Age  int        `json:"age" yaml:"yaml_age"  xiaoyi:"age" `
	Rich bool       `json:"rich" yaml:"-"  xiaoyi:"-" `
	Hc   HttpConfig `yaml:"http"`
}
type HttpConfig struct {
	Ip   []string `yaml:"ips"  `
	Port int      `yaml:"port"  `
}

func jsonwork() {
	p := Person5{
		Name: "leo",
		Age:  18,
		Rich: true,
		Hc:   HttpConfig{},
	}
	data, err := json.Marshal(p)
	if err != nil {
		log.Printf("[json.Marshal.error][err:%v]", err)
		return
	}
	log.Printf("[person.json.Marshal.res:%v]", string(data))
	p2str := `{
	"name":"李逵",
	"age":20,
	"rich":true
}`
	var p2 Person5
	err = json.Unmarshal([]byte(p2str), &p2)
	if err != nil {
		log.Printf("[json.UnMarshal.error][err:%v]", err)
		return
	}
	log.Printf("[person.json.UnMarshal:%v]", p2)
}

func yamlWork() {
	fileName := "a.yaml"
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("[ioutil.ReadFile.error][err:%v]", err)
		return
	}
	var p Person5
	err = yaml.Unmarshal(content, &p)
	if err != nil {
		log.Printf("[yaml.Unmarshal.error][err:%v]", err)
		return
	}
	log.Printf("[yaml.Unmarshal][res:%v]", p)

	p1 := Person5{
		Name: "abc",
		Age:  20,
		Rich: false,
		Hc:   HttpConfig{},
	}
	data, err := yaml.Marshal(p1)
	if err != nil {
		log.Printf("[yaml.Marshal.error][err:%v]", err)
		return
	}
	err = ioutil.WriteFile("b.yaml", data, 0644)
	if err != nil {
		log.Printf("[ioutil.WriteFile.error][err:%v]", err)
		return
	}
}
func structTag(s interface{}) {
	t := reflect.TypeOf(s)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		key := field.Name
		jsonV := field.Tag.Get("json")
		yamlV := field.Tag.Get("yaml")
		if tag, ok := field.Tag.Lookup("xiaoyi"); ok {
			log.Printf("[找到了xiaoyi标签:key:%s xiaoyi:%s]", key, tag)
		}
		log.Printf("[key=%s json=%s yaml=%s]", key, jsonV, yamlV)
	}
}
func myTagWork() {
	p := Person5{
		Name: "123",
		Age:  10,
		Rich: false,
		Hc:   HttpConfig{},
	}
	structTag(p)
}
func main() {
	//jsonwork()
	//yamlWork()
	myTagWork()
}
