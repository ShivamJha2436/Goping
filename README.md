GoPing - A minimal implementation of the classic ping utility written in Go.
It manually constructs ICMP Echo Request packets, sends them over a raw socket, receives replies, and measures RTT. Built for learning low-level networking and seeing the packets “under the hood”.

## 1 — Project purpose & overview
### Why this project exists
ping is one of the simplest and most revealing networking tools. Rebuilding it yourself forces you to learn:
- What ICMP is (a network-layer protocol for diagnostics and errors).
- How to build packets (headers + payload).
- How to calculate the ICMP checksum.
- How to send/receive raw packets from user space (raw sockets).
- How RTT (round trip time) is measured.

### What this project does (current state):
- Resolves a hostname to an IPv4 address.
- Constructs a valid ICMP Echo Request packet (Type=8, Code=0) with an ID, sequence number and a payload.
- Calculates and inserts the ICMP checksum.
- Sends the packet via a raw IPv4 socket (IPPROTO_ICMP).
- Receives a reply (blocking Recvfrom) and measures RTT.
- Prints per-packet RTT output.