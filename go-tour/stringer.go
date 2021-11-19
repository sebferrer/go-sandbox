package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
/*
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}*/

func (ipaddr IPAddr) String() string {
	var strs []string
	for _, v := range ipaddr {
		strs = append(strs, strconv.Itoa(int(v)))
	}
	return strings.Join(strs, ".")
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

/** Output

loopback: 127.0.0.1
googleDNS: 8.8.8.8
*/
