package main

import (
	"fmt"
	"net"
	"os"
	"syscall"
	"unsafe"
)

func getAdapterList() (*syscall.IpAdapterInfo, error) {
	b := make([]byte, 1000)
	l := uint32(len(b))
	a := (*syscall.IpAdapterInfo)(unsafe.Pointer(&b[0]))
	// TODO(mikio): GetAdaptersInfo returns IP_ADAPTER_INFO that
	// contains IPv4 address list only. We should use another API
	// for fetching IPv6 stuff from the kernel.
	err := syscall.GetAdaptersInfo(a, &l)
	if err == syscall.ERROR_BUFFER_OVERFLOW {
		b = make([]byte, l)
		a = (*syscall.IpAdapterInfo)(unsafe.Pointer(&b[0]))
		err = syscall.GetAdaptersInfo(a, &l)
	}
	if err != nil {
		return nil, os.NewSyscallError("GetAdaptersInfo", err)
	}
	return a, nil
}

func localAddresses() error {
	ifaces, err := net.Interfaces()
	if err != nil {
		return err
	}

	aList, err := getAdapterList()
	if err != nil {
		return err
	}

	for _, ifi := range ifaces {
		for ai := aList; ai != nil; ai = ai.Next {
			index := ai.Index

			if ifi.Index == int(index) {
				ipl := &ai.IpAddressList
				for ; ipl != nil; ipl = ipl.Next {

					fmt.Printf("%s: %s (%s)\n", ifi.Name, ipl.IpAddress, ipl.IpMask)
				}
			}
		}
	}
	return err
}

func main() {
	err := localAddresses()
	if err != nil {
		panic(err)
	}
}
