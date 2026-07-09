package dnslookup

import (
	"fmt"
	"net"
)

func RunDnsLookup(isEnabled bool) {
	if !isEnabled {
		return
	}

	fmt.Print("DNS lookup\n\n")

	ips, err := net.LookupHost("api.eventsease.in")
	if err != nil {
		panic(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
