package main

import (
	"fmt"
	"os"
)

func main() {
	userFile := "./TextLearn/File/banming.txt"
	fl, err := os.Open(userFile)
	defer fl.Close()
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}
