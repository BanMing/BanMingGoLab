package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	//if len(os.Args) != 2 {
	//	fmt.Fprintf(os.Stderr, "Usage %s host:port \n", os.Args[0])
	//	os.Exit(1)
	//}
	//service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":7777")
	if err != nil {
		fmt.Println("err:", err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))
}

func checkError(err error) {
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}
}
