package ri

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type Server struct {
	MaxBuffSize int

	clientDB map[string]NetworkInfo
	conn     *net.UDPConn
}

func NewServer() *Server {
	s := new(Server)
	s.clientDB = make(map[string]NetworkInfo)
	s.MaxBuffSize = 2048
	return s
}

//Use ":10001" to listen port 10001
func (s *Server) ListenAndServe(port string) {
	ServerAddr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		log.Println("Server port init error:", err)
		return
	}

	s.conn, err = net.ListenUDP("udp", ServerAddr)
	if err != nil {
		log.Println("Server listening  error:", err)
		return
	}
	defer s.conn.Close()
	log.Println("UDP Server strating listen:", port)
	buf := make([]byte, s.MaxBuffSize)

	for {
		n, addr, err := s.conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			s.parseCmds(string(buf[0:n]), addr.String())
		}
	}

}

func (s *Server) routingInfo(cmd, addr string) {
	ni, err := DecodeRoutingInfo(cmd)
	if err != nil {
		log.Println("DecodeRoutingInfo err=", err)
		return
	}

	ip, port := DecodeIpPort(addr)
	ni.EIPv4 = ip
	ni.EPort = port
	s.clientDB[ni.Id] = *ni
	log.Println("RoutingInfo work:", ni, " is it use NAT?", ni.UseNAT())
}

func (s *Server) parseCmds(cmd string, addr string) {
	fmt.Println("Received ", cmd, " from ", addr)
	if strings.Contains(cmd, CMD_RoutingInfo) {
		log.Println("Cmd:", CMD_RoutingInfo)
		s.routingInfo(cmd, addr)

	} else if strings.Contains(cmd, CMD_HolePunching) {
		log.Println("Cmd:", CMD_HolePunching)

	} else if strings.Contains(cmd, CMD_RequestPairing) {
		log.Println("Cmd:", CMD_RequestPairing)

	} else {
		log.Println("Cmd invalid.")
	}

}
