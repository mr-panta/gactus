package util

import "net"

// GetIPAddrs is used to get ip version 4 addressses
func GetIPAddrs() (ipAddrs []string, err error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip.To4() != nil {
				ipAddrs = append(ipAddrs, ip.String())
			}
		}
	}
	return ipAddrs, nil
}
