package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	fmt.Println("URL Parsing")

	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Scheme:", u.Scheme)
	fmt.Println("User:", u.User)
	fmt.Println("UserName:", u.User.Username())
	p, isPresent := u.User.Password()
	fmt.Println("Password:", p, isPresent)

	fmt.Println("Host:", u.Host)
	host, port, err := net.SplitHostPort(u.Host)
	if err != nil {
		panic(err)
	}
	fmt.Println("Host:", host, "Port:", port)

	fmt.Println("Path:", u.Path)
	fmt.Println("Fragment:", u.Fragment)

	fmt.Println("RawQuery:", u.RawQuery)
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println("ParseQuery:", m)
	fmt.Println("ParseQuery:", m["k"][0])
}
