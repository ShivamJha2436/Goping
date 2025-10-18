package icmp

// Checksum calculates the standard Internet checksum (RFC 1071)
func Checksum(data []byte) uint16 {
	var sum uint32

	// Sum 16-bit words
	for i := 0; i < len(data)-1; i += 2 {
		sum += uint32(data[i])<<8 + uint32(data[i+1])
	}

	// If odd number of bytes, add last byte padded with zero
	if len(data)%2 == 1 {
		sum += uint32(data[len(data)-1]) << 8
	}

	// Fold sum to 16 bits
	for (sum >> 16) > 0 {
		sum = (sum & 0xffff) + (sum >> 16)
	}

	// One’s complement
	return ^uint16(sum)
}
