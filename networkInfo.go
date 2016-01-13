package ri

import (
	"log"
	"net"
	"os/exec"
	"runtime"
	"strings"
)

type NetworkInfo struct {
	EIPv4 string
	EIPv6 string
	EPort int

	IIPv4    string
	IIPv6    string
	IPort    int
	INetmask string
}

func NewNetworkInfo() *NetworkInfo {
	ni := new(NetworkInfo)
	return ni
}

func (n *NetworkInfo) InitInternalInfo() {

}

func (n *NetworkInfo) CheckRouting() bool {
	return false
}

func GetNetworkMask(deviceName string) string {
	switch runtime.GOOS {
	case "darwin":
		cmd := exec.Command("ipconfig", "getoption", deviceName, "subnet_mask")
		out, err := cmd.CombinedOutput()
		if err != nil {
			return ""
		}

		nm := strings.Replace(string(out), "\n", "", -1)
		log.Println("netmask=", nm, " OS=", runtime.GOOS)
		return nm
	default:
		return ""
	}
	return ""
}

func GetIPv4Mask() (net.IP, string) {

	ifaces, err := net.Interfaces()
	// handle err
	if err != nil {
		log.Println("No network:", err)
		return nil, ""
	}

	for _, i := range ifaces {
		if strings.Contains(i.Name, "en") {
			addrs, err := i.Addrs()
			// handle err
			if err != nil {
				log.Println("No IP:", err)
				return nil, ""
			}

			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					log.Println("IPNET")
					ip = v.IP
				case *net.IPAddr:
					log.Println("IPAddr")
					ip = v.IP
				}
				if ip[0] == 0 {
					log.Println("Get device:", i.Name)
					return ip, GetNetworkMask(i.Name)
				}
			}
		}
	}

	return nil, ""
}

func FindInternalPort(conn *net.UDPConn) string {
	if len(conn.LocalAddr().String()) <= 0 {
		return 0
	}

	log.Println("network:", conn.LocalAddr().Network(), " add:", conn.LocalAddr().String())
	iport := conn.LocalAddr().String()[strings.Index(conn.LocalAddr().String(), ":")+1:]
	log.Println("port:", iport)
	return iport
}
