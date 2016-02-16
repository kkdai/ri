package main

import (
	"flag"
	"fmt"

	. "github.com/kkdai/ri"
)

func main() {
	flag.Parse()
	cmds := flag.Args()
	ser := NewServer()
	riCallback := func(ni *NetworkInfo, err error) (retErr error) {
		//Your callback here
		fmt.Println("Got RI:", ni, " err=", err)
		return nil
	}

	ser.ListenAndServe(cmds[0], riCallback)
}
