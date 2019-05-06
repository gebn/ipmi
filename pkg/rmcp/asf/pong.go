package asf

import (
	"encoding/binary"
	"fmt"
)

// PresencePong defines the structure of a Presence Pong message's payload.
type PresencePong struct {

	// Enterprise is the IANA Enterprise Number of an entity that has defined
	// OEM-specific capabilities for the managed client. If no such capabilities
	// exist, this is set to ASF's IANA Enterprise Number.
	Enterprise uint32

	// OEM identifies OEM-specific capabilities. Its structure is defined by the
	// OAM. This is set to 0s if no OEM-specific capabilities exist. This
	// implementation does not change byte order from the wire for this field.
	OEM [4]byte

	// Entities defines supported entities. In practice, there is a bit for an
	// unspecified version of IPMI (v1 at the time of the RMCP spec), and a bit
	// for ASF-RMCP v1.0 itself.
	Entities uint8

	// 7 bytes reserved, set to 0s.
}

// Parse makes the struct represent the specified payload.
func (p *PresencePong) Parse(b []byte) error {
	if len(b) != 16 {
		return fmt.Errorf("payload must be 16 bytes long, got %v", len(b))
	}

	p.Enterprise = binary.BigEndian.Uint32(b[0:4])
	copy(p.OEM[:], b[4:8]) // N.B. no byte order change
	p.Entities = uint8(b[8])
	// ignore remaining 7 bytes; should be set to 0s
	return nil
}

// String satisfies Stringer for PresencePong.
func (p *PresencePong) String() string {
	return fmt.Sprintf("PresencePong(Enterprise: %v, SupportsIPMI: %v)",
		p.Enterprise, p.SupportsIPMI())
}

// SupportsIPMI returns whether the Presence Pong message indicates support for
// IPMIv1.
func (p *PresencePong) SupportsIPMI() bool {
	return p.Entities&(1<<7) != 0
}

// SupportsASFV1 returns whether the Presence Pong message indicates support for
// ASF, Version 1.0. This seems somewhat redundant, given the request is made
// using ASF v1.0.
func (p *PresencePong) SupportsASFV1() bool {
	return p.Entities&1 != 0
}
