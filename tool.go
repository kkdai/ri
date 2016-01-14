package ri

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type NATType int

const (
	NAT_NotAvailiable  NATType = iota
	NAT_FullCone       NATType = iota
	NAT_IPRestricted   NATType = iota
	NAT_PortRestricted NATType = iota
	NAT_Symmetric      NATType = iota
)

const (
	CMD_RoutingInfo    string = "RoutingInformation"
	CMD_RequestPairing string = "RequestPairing"
	CMD_HolePunching   string = "HolePunching"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
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
		//return 255.255.255.0 as default
		return "255.255.255.0"
	}
	return "255.255.255.0"
}

func EncodeRoutingInfo(id string, n *ClientNI) string {
	if n == nil {
		return ""
	}

	return fmt.Sprintf("%s %s,%s,%s,%d,%s", CMD_RoutingInfo, id, n.IIPv4, n.IIPv6, n.IPort, n.INetmask)
}
