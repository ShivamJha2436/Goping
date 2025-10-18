package net

import (
	"fmt"
	"net"
	"syscall"
	"time"
)

// SendICMP sends an ICMP packet to the given destination IP
func SendICMP(destIP string, packet []byte, timeout time.Duration) (time.Duration, error) {
	// Create a raw socket (IPv4, ICMP)
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	if err != nil {
		return 0, fmt.Errorf("failed to create raw socket: %v", err)
	}
	defer syscall.Close(fd)

	// Set a receive timeout (so we don't block forever)
	tv := syscall.NsecToTimeval(timeout.Nanoseconds())
	if err := syscall.SetsockoptTimeval(fd, syscall.SOL_SOCKET, syscall.SO_RCVTIMEO, &tv); err != nil {
		return 0, fmt.Errorf("failed to set timeout: %v", err)
	}

	// Resolve destination IP to syscall.Sockaddr
	dst := &syscall.SockaddrInet4{}
	ip := net.ParseIP(destIP).To4()
	copy(dst.Addr[:], ip)

	// Record send time
	start := time.Now()

	// Send ICMP Echo Request
	if err := syscall.Sendto(fd, packet, 0, dst); err != nil {
		return 0, fmt.Errorf("failed to send packet: %v", err)
	}

	// Buffer to receive reply
	reply := make([]byte, 1500)
	_, _, err = syscall.Recvfrom(fd, reply, 0)
	if err != nil {
		return 0, fmt.Errorf("failed to receive reply: %v", err)
	}

	rtt := time.Since(start)
	return rtt, nil
}
