package main

import (
	"fmt"
	"net"
)

func main() {
	/* Lets prepare a address at any address at port 10001*/
	ServerAddr, _ := net.ResolveUDPAddr("udp", ":10001")

	/* Now listen at selected port */
	ServerConn, _ := net.ListenUDP("udp", ServerAddr)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		fmt.Println("Received ", string(buf[0:n]), " from ", addr)

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
