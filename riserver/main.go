package main

import (
	"flag"

	. "github.com/kkdai/ri"
)

func main() {
	flag.Parse()
	cmds := flag.Args()
	ser := NewServer()
	ser.ListenAndServe(cmds[0])
}
