package main

import (
	"fmt"
	"os"
)

func main() {
	os.Mkdir("./banming", 0777)
	os.MkdirAll("./banming/test1/test2", 0777)
	err := os.Remove("./banming")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("./banming")
}
