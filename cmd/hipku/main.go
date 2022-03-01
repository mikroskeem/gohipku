package main

import (
	"fmt"
	"net"
	"os"

	"github.com/mikroskeem/gohipku/pkg"
)

func main() {
	ipRaw := os.Args[1]
	ip := net.ParseIP(ipRaw)

	value, ok := gohipku.Encode(ip)
	if !ok {
		fmt.Fprintln(os.Stderr, "invalid ip")
		os.Exit(1)
	}

	fmt.Println(value)
}
