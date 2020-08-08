package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {
	var ip string
	flag.StringVar(&ip, "ip", "", "123.123.123.123")
	flag.Parse()

	if ip == "" {
		flag.PrintDefaults()
		log.Fatal("IP not set")
	}

	fmt.Println(isCloudflare(ip))
}

var cloudflareCIDRs = [14]string{
	"173.245.48.0/20",
	"103.21.244.0/22",
	"103.22.200.0/22",
	"103.31.4.0/22",
	"141.101.64.0/18",
	"108.162.192.0/18",
	"190.93.240.0/20",
	"188.114.96.0/20",
	"197.234.240.0/22",
	"198.41.128.0/17",
	"162.158.0.0/15",
	"104.16.0.0/12",
	"172.64.0.0/13",
	"131.0.72.0/22",
}

func isCloudflare(ip string) bool {
	for _, CIDR := range cloudflareCIDRs {
		_, ipnetA, _ := net.ParseCIDR(CIDR)
		ipAddr, _, _ := net.ParseCIDR(ip + "/0")

		if ipnetA.Contains(ipAddr) {
			return true
		}
	}
	return false
}
