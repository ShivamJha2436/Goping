package icmp

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"time"
)

// ICMP represents an ICMP Echo Request/Reply packet
type ICMP struct {
	Type     uint8
	Code     uint8
	Checksum uint16
	ID       uint16
	Seq      uint16
	Data     []byte
}

// NewICMPEcho creates a new ICMP Echo Request packet
func NewICMPEcho(seq uint16) *ICMP {
	rand.Seed(time.Now().UnixNano())

	return &ICMP{
		Type: 8, // Echo Request
		Code: 0,
		ID:   uint16(rand.Intn(0xffff)),
		Seq:  seq,
		Data: []byte("hello from Shivam's Go ping!"),
	}
}

// Marshal converts ICMP struct → bytes, with checksum filled
func (p *ICMP) Marshal() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Write header with 0 checksum first
	binary.Write(buf, binary.BigEndian, p.Type)
	binary.Write(buf, binary.BigEndian, p.Code)
	binary.Write(buf, binary.BigEndian, uint16(0))
	binary.Write(buf, binary.BigEndian, p.ID)
	binary.Write(buf, binary.BigEndian, p.Seq)
	buf.Write(p.Data)

	raw := buf.Bytes()
	checksum := Checksum(raw)

	// Rewrite with correct checksum
	buf = new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, p.Type)
	binary.Write(buf, binary.BigEndian, p.Code)
	binary.Write(buf, binary.BigEndian, checksum)
	binary.Write(buf, binary.BigEndian, p.ID)
	binary.Write(buf, binary.BigEndian, p.Seq)
	buf.Write(p.Data)

	return buf.Bytes(), nil
}
