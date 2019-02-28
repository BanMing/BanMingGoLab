package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
)

func main() {

	js, err := simplejson.NewJson([]byte(
		`
		{"test":{
		"array":[1,"2",3],
		"int":10,
		"float":5.102,
		"bignum":999999999999999999,
		"string":"simplejson",
		"bool":true
	}}`))
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("@")
	fmt.Println(js.Get("test").Get("int").Int())
}
