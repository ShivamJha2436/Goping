GoPing - A minimal implementation of the classic ping utility written in Go.
It manually constructs ICMP Echo Request packets, sends them over a raw socket, receives replies, and measures RTT. Built for learning low-level networking and seeing the packets “under the hood”.

## 1 — Project purpose & overview
### Why this project exists
`ping` is one of the simplest and most revealing networking tools. Rebuilding it yourself forces you to learn:
- What ICMP is (a network-layer protocol for diagnostics and errors).
- How to build packets (headers + payload).
- How to calculate the ICMP checksum.
- How to send/receive raw packets from user space (raw sockets).
- How RTT (round trip time) is measured.

### What this project does:
- Resolves a hostname to an IPv4 address.
- Constructs a valid ICMP Echo Request packet (Type=8, Code=0) with an ID, sequence number and a payload.
- Calculates and inserts the ICMP checksum.
- Sends the packet via a raw IPv4 socket (IPPROTO_ICMP).
- Receives a reply (blocking Recvfrom) and measures RTT.
- Prints per-packet RTT output.

## 2 — Networking background:
### ICMP
- Stands for Internet Control Message Protocol.
- Used for network diagnostics and error messages.
- ping uses ICMP Echo Request (Type = 8) and Echo Reply (Type = 0).
- ICMP is carried directly by IP (protocol number 1).

### Raw sockets
- Raw sockets let you send and receive network packets with minimal kernel assistance (you supply the payload).
- Creating a raw socket requires elevated privileges (root) on most OSes.
- For IPv4 ICMP we use: socket(AF_INET, SOCK_RAW, IPPROTO_ICMP).

### ICMP packet layout (Echo Request/Reply)
```sh
0         1         2         3
+---------+---------+---------+---------+
| Type(1) | Code(1) | Checksum(2)       |
+---------------------------------------+
| Identifier(2) | Sequence Number(2)    |
+---------------------------------------+
| Data (variable)                       |
+---------------------------------------+
```
### Internet checksum (RFC 1071)
- 16-bit ones-complement sum of all 16-bit words in the header+data, with odd byte padded.
- Final checksum = ones-complement of the folded sum.