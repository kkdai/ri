package main

import (
	"flag"

	. "github.com/kkdai/ri"
)

func main() {
	flag.Parse()
	serAdd := flag.Args()

	c := NewClient()
	c.Id = "test/1234"
	c.ConnectTo(serAdd[0])
	c.SendRoutingInfo()
}
