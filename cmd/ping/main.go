package main

import (
	"fmt"
	"os"

	internal "github.com/ShivamJha2436/ping/internal"
	myNet "github.com/ShivamJha2436/ping/internal/net"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ping <hostname>")
		os.Exit(1)
	}

	host := os.Args[1]
	ip, err := myNet.ResolveHost(host)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	internal.RunPing(host, ip, 4)
}
