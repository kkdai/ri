package ri

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
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
}

func DecodeIpPort(localAddr string) (string, int) {
	if len(localAddr) <= 0 {
		return "", 0
	}

	ip := localAddr[:strings.Index(localAddr, ":")]
	log.Println("ip:", ip)

	//Get Port
	iport := localAddr[strings.Index(localAddr, ":")+1:]
	log.Println("port:", iport)
	nPort, err := strconv.Atoi(iport)
	if err != nil {
		return "", 0
	}

	return ip, nPort
}

func EncodeRoutingInfo(id string, n *NetworkInfo) string {
	if n == nil {
		return ""
	}

	return fmt.Sprintf("%s %s,%s,%s,%d,%s", CMD_RoutingInfo, id, n.IIPv4, n.IIPv6, n.IPort, n.INetmask)
}

func DecodeRoutingInfo(ri string) (*NetworkInfo, error) {
	log.Println("DecodeRoutingInfo:", ri)
	if len(ri) == 0 || !strings.Contains(ri, CMD_RoutingInfo) {
		return nil, errors.New("Invalid Input.")
	}

	var id, iip4, iip6, imask string
	var iport int
	ri = strings.Replace(ri, ",", " ", -1)
	n, err := fmt.Sscanf(ri, CMD_RoutingInfo+" %s %s %s %d %s", &id, &iip4, &iip6, &iport, &imask)
	if err != nil {
		log.Println(n, err)
		return nil, err
	}

	log.Println("Got:", n, "=>", id, iip4, iip6, imask, iport)
	ni := new(NetworkInfo)
	ni.IIPv4 = iip4
	ni.IIPv6 = iip6
	ni.Id = id
	ni.IPort = iport
	ni.INetmask = imask
	return ni, nil
}
