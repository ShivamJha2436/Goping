package internal

import (
	"fmt"
	"time"

	"github.com/ShivamJha2436/ping/internal/icmp"
	myNet "github.com/ShivamJha2436/ping/internal/net"
)

func RunPing(host, ip string, count int) {
	fmt.Printf("PING %s (%s):\n", host, ip)

	for i := 1; i <= count; i++ {
		pkt := icmp.NewICMPEcho(uint16(i))
		raw, _ := pkt.Marshal()

		rtt, err := myNet.SendICMP(ip, raw, 2*time.Second)
		if err != nil {
			fmt.Printf("Request timeout for icmp_seq=%d\n", i)
			continue
		}

		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%.2f ms\n",
			len(raw), ip, i, float64(rtt.Microseconds())/1000.0)

		time.Sleep(1 * time.Second)
	}
}
