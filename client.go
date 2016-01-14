package ri

import (
	"fmt"
	"log"
	"net"
	"time"
)

type Client struct {
	Id   string
	conn *net.UDPConn
	ni   *ClientNI
	//UDP resend time to ensure UDP socket lost
	UDPResend int
}

func NewClient() *Client {
	c := new(Client)
	c.UDPResend = 8
	return c
}

func (c *Client) ConnectTo(srvAddr string) error {
	ServerAddr, err := net.ResolveUDPAddr("udp", srvAddr)
	if err != nil {
		return err
	}

	Conn, err := net.DialUDP("udp", nil, ServerAddr)
	if err != nil {
		return err
	}

	c.conn = Conn

	c.ni = NewClientNI()
	c.ni.InitNetworkInfo(c.conn.LocalAddr().String())
	return nil
}

func (c *Client) SendRoutingInfo() error {
	for i := 0; i < c.UDPResend; i++ {
		msg := EncodeRoutingInfo(c.Id, c.ni)
		buf := []byte(msg)
		log.Println("write->", msg)
		_, err := c.conn.Write(buf)
		if err != nil {
			fmt.Println(msg, err)
			return err
		}
		time.Sleep(time.Second * 1)
	}
	log.Println("SendRoutingInfo done")
	return nil
}
