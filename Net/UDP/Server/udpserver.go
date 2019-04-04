package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	server := ":7777"
	udpAddr, err := net.ResolveUDPAddr("udp4", server)
	checkError(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)
	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	var buf [512]byte
	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}
	dateTime := time.Now().String()
	fmt.Println("dateTime:", dateTime)
	conn.WriteToUDP([]byte(dateTime), addr)
}
func checkError(err error) {
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}
}
