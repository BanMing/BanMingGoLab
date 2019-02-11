package main

import (
	"fmt"
	"runtime"
)

func sayHi(s string) {
	for i := 1; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	//开启一个协成
	go sayHi("World")
	sayHi("Hello")
}
