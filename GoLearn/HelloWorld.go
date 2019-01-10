package main

import "fmt"

func main() {
	// fmt.Printf("Hello world")
	FixStr()
}

var frenchHello string
var emptyString string = ""

// 修改字符串
func FixStr() {
	var s string = "hello"
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Println(s2)
}

//切片字符串
func SpitStr() {
	s := "hello"
	c := "c" + s[1:]
	fmt.Println(c)

}
