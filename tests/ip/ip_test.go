package ip

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"testing"
)

func TestIp(t *testing.T) {
	ip := net.ParseIP("192.168.1.20")
	fmt.Println(ip.DefaultMask().String())

	fmt.Println(StringIpToUint("192.168.1.20"))
	ips := StringIpToUint("192.168.1.20")

	fmt.Println(UintIpToString(ips))

	fmt.Println(0b0011 | 0b1100)
	fmt.Println(0b1111)
}

func StringIpToUint(ip string) (uip uint32) {
	pos := 24
	ips := strings.Split(ip, ".")
	for _, item := range ips {
		v, _ := strconv.Atoi(item)
		v = v << pos
		uip = uip | uint32(v)
		pos -= 8
	}
	return
}

func UintIpToString(ip uint32) string {
	ipUint := []uint32{ip >> 24 & 0xFF, ip >> 16 & 0xFF, ip >> 8 & 0xFF, ip & 0xFF}
	var ips []string
	for _, item := range ipUint {
		s := strconv.Itoa(int(item))
		ips = append(ips, s)
	}
	return strings.Join(ips, ".")
}
