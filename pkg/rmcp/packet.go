package rmcp

import (
	"fmt"
)

// Packet describes the format of an RMCP packet, which forms a UDP
// payload. See section 3.2.2.2.
type Packet struct {

	// Version identifies the version of the RMCP header.
	// 0x06 indicates RMCP v1.0; lower values are legacy, higher
	// values are reserved.
	Version uint8

	// 1 byte reserved, should be set to 0x00.

	// Sequence is the sequence number assicated with the message.
	// Note that this rolls over to 0 after 254, not 255. Seq num
	// 255 indicates the receiver must not send an ACK.
	Sequence uint8

	// Class idicates whether this packet is an ACK, and the
	// structure of the payload.
	Class uint8

	// Data contains the payload of the packet. It will be empty
	// if the packet is an ACK.
	Data []byte
}

// IsAck returns true if the RMCP packet is an acknowledgment, or false
// if it is a "normal" packet. See section 3.2.2.1.
func (p *Packet) IsAck() bool {
	return p.Class&Ack == Ack
}

// Marshal transforms the packet into its wire format.
func (p *Packet) Marshal() []byte {
	m := make([]byte, 1+1+1+1+len(p.Data))
	m[0] = p.Version
	m[2] = p.Sequence
	m[3] = p.Class
	copy(m[4:], p.Data)
	return m
}

// Parse modifies the packet in-place to represent a wire format.
func (p *Packet) Parse(b []byte) error {
	if len(b) < 4 {
		return fmt.Errorf("an RMCP packet must be at least 4 bytes for the header")
	}

	p.Version = uint8(b[0])
	// 1 byte reserved
	p.Sequence = uint8(b[2])
	p.Class = uint8(b[3])
	p.Data = b[4:]
	return nil
}

func (p *Packet) String() string {
	return fmt.Sprintf("Packet(Version: %#v, Seq: %v, Ack: %v, Class: %#v, Data: %v)",
		p.Version, p.Sequence, p.IsAck(), p.Class&0xF, p.Data)
}
