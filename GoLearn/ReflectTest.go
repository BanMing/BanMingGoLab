package main

import (
	//"GoLearn.CommaTest.main"
	"fmt"
	"reflect"
)

func main() {
	i := Person{"testName", 2}
	t := reflect.TypeOf(i)             // 获得类型的元数据，通过t我们能获取类型定义里面的所有元素
	v := reflect.ValueOf(i)            //获得实际的值，通过v我们可以获得存储在里面的值，还可以去改变值
	tag := t.Elem().Field(0).Tag       //获取定义struct里面的标签
	name := v.Elem().Field(0).String() //获取存储在第一个字段的值
	fmt.Println(v)
	fmt.Println(tag)
	fmt.Println(name)
}
