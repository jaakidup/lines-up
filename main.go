package main

import (
	"errors"
	"fmt"
	"log"
	"net"
)

func main() {
	ipAddresses, err := getExternalIPAddresses()
	checkError(err)
	for _, ip := range ipAddresses {
		fmt.Println(ip)
	}

}

func getExternalIPAddresses() ([]string, error) {
	var ipAddresses []string

	interfaces, err := net.Interfaces()
	checkError(err)

	for _, interf := range interfaces {

		if interf.Flags&net.FlagUp == 0 {
			continue
		}

		if interf.Flags&net.FlagLoopback != 0 {
			continue
		}

		addresses, err := interf.Addrs()
		checkError(err)
		for _, address := range addresses {

			var ip net.IP

			switch a := address.(type) {
			case *net.IPNet:
				ip = a.IP
			case *net.IPAddr:
				ip = a.IP
			}

			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not ipv4
			}
			ipAddresses = append(ipAddresses, ip.String())
		}
	}
	if len(ipAddresses) == 0 {
		return ipAddresses, errors.New("No interfaces up or connected")
	}
	return ipAddresses, nil
}

func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}
