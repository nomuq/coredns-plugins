package main

import (
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/coremain"

	_ "github.com/coredns/coredns/plugin/forward"
	_ "github.com/coredns/coredns/plugin/log"
)

func init() {
	dnsserver.Directives = []string{
		"log",
		"forward",
	}
}

func main() {
	coremain.Run()
}
