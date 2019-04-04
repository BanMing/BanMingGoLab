package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	tcpServer()
}

//high performance
func tcpServer() {
	servive := ":7777"
	//resolve ip to address
	tcpAddr, err := net.ResolveTCPAddr("tcp4", servive)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	dateTime := time.Now().String()
	conn.Write([]byte(dateTime))
}

func tcpServer1() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		conn.Close()
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}
}
