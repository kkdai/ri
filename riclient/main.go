package main

import (
	"flag"

	. "github.com/kkdai/ri"
)

func main() {
	flag.Parse()
	cmds := flag.Args()

	c := NewClient()
	c.Id = "test/1234"
	c.ConnectTo(cmds[0])
	c.SendRoutingInfo()
}
