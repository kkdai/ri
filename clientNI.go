package ri

import (
	"errors"
	"log"
	"net"
	"strconv"
	"strings"
)

type ClientNI struct {
	IIPv4    string
	IIPv6    string
	IPort    int
	INetmask string
}

func NewClientNI() *ClientNI {
	ni := new(ClientNI)
	return ni
}

func (n *ClientNI) InitNetworkInfo(localAddr string) error {
	err := n.getInternalPort(localAddr)
	if err != nil {
		return err
	}

	err = n.enumDevice()
	return err
}

func (n *ClientNI) enumDevice() error {

	ifaces, err := net.Interfaces()
	// handle err
	if err != nil {
		log.Println("No network:", err)
		return err
	}

	for _, i := range ifaces {
		if !strings.Contains(i.Name, "en") {
			continue
		}

		addrs, err := i.Addrs()
		// handle err
		if err != nil {
			log.Println("No IP:", err)
			continue
		}

		var ipv6 string
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			locateAdapater := false

			if ip[0] != 0 {
				ipv6 = ip.String()
			} else {
				//IPv4
				n.INetmask = GetNetworkMask(i.Name)
				n.IIPv4 = ip.String()
				locateAdapater = true
				log.Println("Find ipv4 mapping:", ip.String(), n.INetmask, i.Name)
			}

			if locateAdapater {
				n.IIPv6 = ipv6
				return nil
			}
		}
	}

	return errors.New("Not find specific IP")
}

func (n *ClientNI) getInternalPort(localAddr string) error {
	if len(localAddr) <= 0 {
		return errors.New("No exist UDP connection.")
	}

	//Get Port
	iport := localAddr[strings.Index(localAddr, ":")+1:]
	log.Println("port:", iport)
	nPort, err := strconv.Atoi(iport)
	if err != nil {
		return err
	}

	n.IPort = nPort
	return nil
}
