package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	p, _ := url.Parse(s)

	fmt.Println(p.Scheme)

	fmt.Println(p.User.Username())
	fmt.Println(p.User.Password())
	fmt.Println(p.Host)

	host, port, _ := net.SplitHostPort(p.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(url.ParseQuery(p.RawQuery))
}
