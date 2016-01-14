RI: A simple routing information client/server
==================

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/kkdai/ri/master/LICENSE)  [![GoDoc](https://godoc.org/github.com/kkdai/ri?status.svg)](https://godoc.org/github.com/kkdai/ri)  [![Build Status](https://travis-ci.org/kkdai/ri.svg?branch=master)](https://travis-ci.org/kkdai/ri)
    

**RI** is a UDP client/server to get Public and Private IP and Port for firewall penetration

It could get your IPv4 and IPv6 information during client/server communication. It also support get the IPv4 Network Mask.

It could get public IP and Port, if your server host in a public IP. It will be prepare stage for next step of firewall [Hole Punching](https://en.wikipedia.org/wiki/Hole_punching_(networking)).

Install
---------------
Install package `go get github.com/kkdai/ri`

Install server: `go get github.com/kkdai/riserver`

Install client: `go get github.com/kkdai/riclient`


Usage
---------------

#### Server side example

Init a local UDP server on port "10001".

```go
package main

import . "github.com/kkdai/ri"

func main() {
	ser := NewServer()
	ser.ListenAndServe(":10001")
}

```

#### Client side example

Send UDP socket to server "127.0.0.1:10001"

```go
package main

import (
	"flag"

	. "github.com/kkdai/ri"
)

func main() {
	c := NewClient()
	c.Id = "test/1234"
	c.ConnectTo("127.0.0.1:10001")
	c.SendRoutingInfo()
}

```

###Use the binary directly

- Init server with port "10001"
	- `riserver ":10001"`
- Send to server on "192.168.1.1" port "10001"
	- `riclient "192.168.1.1:10001"`

Inspired
---------------

- [GoDoc: net](https://golang.org/pkg/net/)


Project52
---------------

It is one of my [project 52](https://github.com/kkdai/project52).


License
---------------

This package is licensed under MIT license. See LICENSE for details.

