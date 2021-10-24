package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (t IPAddr) String() string {
	res := ""
	dot := "."
	for i := range t {
		if i != len(t)-1 {
			res += strconv.Itoa(int(t[i])) + dot
			continue
		}
		res += strconv.Itoa(int(t[i]))
	}
	return res
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
