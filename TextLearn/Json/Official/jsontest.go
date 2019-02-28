package main

import (
	"encoding/json"
	"fmt"
)

type JsonSever struct {
	ServerName string
	ServerIP   string
}
type Serverslice struct {
	Servers []JsonSever
}

func main() {
	//var s Serverslice
	//大小写不敏感
	//除非是FOO或是FoO
	//赋值字段一定是大写，不然会被忽略
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	//赋值给指定数据结构
	//json.Unmarshal([]byte(str), &s)
	var f interface{}
	err := json.Unmarshal([]byte(str), &f)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(f[""])
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string ", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array: ")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don`t know how to handle")
		}
	}
}
