package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func GetNetMask(deviceName string) string {
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

func GetIP() (net.IP, string) {

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
					return ip, GetNetMask(i.Name)
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

func main() {
	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:10001")
	CheckError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	CheckError(err)

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)
	log.Println("network:", Conn.LocalAddr().Network(), " add:", Conn.LocalAddr().String())
	iport := Conn.LocalAddr().String()[strings.Index(Conn.LocalAddr().String(), ":")+1:]
	log.Println("port:", iport)

	defer Conn.Close()

	for {
		ip, mask := GetIP()
		log.Println("ip=", ip, ip.To4(), ip.DefaultMask(), mask)
		msg := fmt.Sprintf("ip:%s/mask:%s", ip.To4().String(), mask)
		buf := []byte(msg)
		_, err := Conn.Write(buf)
		if err != nil {
			fmt.Println(msg, err)
		}
		time.Sleep(time.Second * 1)
	}
}
