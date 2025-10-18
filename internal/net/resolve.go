package net

import (
	"fmt"
	"net"
)

// ResolveHost resolves a hostname (like google.com) to its IP address
func ResolveHost(host string) (string, error) {
	ipAddr, err := net.ResolveIPAddr("ip4", host)
	if err != nil {
		return "", fmt.Errorf("failed to resolve host %s: %v", host, err)
	}
	return ipAddr.IP.String(), nil
}
