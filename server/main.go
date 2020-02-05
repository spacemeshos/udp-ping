package main

import (
	"fmt"
	"net"
	"strconv"
)

/*
func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("spacemesh"), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}*/

func main() {
	p := make([]byte, 2048)
	port := 7555

	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP("[::]"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	println("Listening on udp port " + strconv.Itoa(port) + "...")

	for {
		_, remoteaddr, err := ser.ReadFromUDP(p)
		fmt.Printf("Message from %v %s \n", remoteaddr, p)
		if err != nil {
			fmt.Printf("Error: %v", err)
			continue
		}
		// go sendResponse(ser, remoteaddr)
	}
}
