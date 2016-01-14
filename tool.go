package ri

import (
	"errors"
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

func DecodeRoutingInfo(ri string) error {
	if len(ri) == 0 || !strings.Contains(ri, CMD_RequestPairing) {
		return errors.New("Invalid Input.")
	}

	var id, resource, iip4, iip6, imask string
	var iport int
	ri = strings.Replace(ri, ",", " ", -1)
	n, err := fmt.Sscanf(ri, CMD_RequestPairing+" %s %s %s %d %s", &id, &iip4, &iip6, &iport, &imask)
	if err != nil {
		fmt.Println(n, err, id, "-", resource)
		return err
	}

	fmt.Println("Got:", n, "=>", id, iip4, iip6, imask, iport)
	return nil
}
