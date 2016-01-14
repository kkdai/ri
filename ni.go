package ri

import (
	"errors"
	"log"
	"net"
	"strings"
)

type NetworkInfo struct {
	Id string

	EIPv4 string
	EPort int

	IIPv4    string
	IIPv6    string
	IPort    int
	INetmask string
}

func NewNI() *NetworkInfo {
	ni := new(NetworkInfo)
	return ni
}

//Return the NAT type according its Internal IP and Public IP
func (n *NetworkInfo) UseNAT() bool {
	return n.IIPv4 == n.EIPv4 && n.IPort == n.EPort
}

//Check each NAT type and make sure if it is under the same routing.
//Return "true" if it is not NAT_Symmetric
func (n *NetworkInfo) ValidToP2P(si *NetworkInfo) bool {
	return si != nil
}
func (n *NetworkInfo) InitNetworkInfo(localAddr string) error {
	ip, port := DecodeIpPort(localAddr)
	if len(ip) != 0 {
		n.IPort = port
	}

	err := n.enumDevice()
	return err
}

func (n *NetworkInfo) enumDevice() error {

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
