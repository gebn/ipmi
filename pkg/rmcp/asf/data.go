package asf

import (
	"encoding/binary"
	"fmt"
)

// Data defines the generic ASF-RMCP message Data block format. See section
// 3.2.2.3 in
// https://www.dmtf.org/sites/default/files/standards/documents/DSP0114.pdf for
// full details.
type Data struct {

	// Enterprise is the IANA Enterprise Number associated with the entity that
	// defines the message type. A list can be found at
	// https://www.iana.org/assignments/enterprise-numbers/enterprise-numbers.
	// This can be thought of as the namespace for the message type. N.B.
	// network byte order.
	Enterprise uint32

	// Type is the message type, defined by the entity associated with the
	// enterprise above.
	Type uint8

	// Tag is the message tag, used to match request-response pairs. The tag of
	// a response is set to that of the message it is responding to. If a
	// message is not expecting a response, this is set to 255.
	Tag uint8

	// 1 byte reserved, set to 0x00.

	// Length is the length of the payload in bytes.
	Length uint8

	// Data is structured as defined by the enterprise and message type.
	Data []byte
}

// Marshal transforms the data into its wire format.
func (d *Data) Marshal() []byte {
	m := make([]byte, 4+1+1+1+1+len(d.Data))
	binary.BigEndian.PutUint32(m[0:4], d.Enterprise)
	m[4] = d.Type
	m[5] = d.Tag
	// 1 byte reserved
	m[7] = d.Length
	copy(m[8:], d.Data)
	return m
}

// Parse modifies the message in-place to represent a wire format.
func (d *Data) Parse(b []byte) error {
	if len(b) < 8 {
		return fmt.Errorf("length must be at least 8 bytes, got %v", len(b))
	}

	d.Enterprise = binary.BigEndian.Uint32(b[0:4])
	d.Type = uint8(b[4])
	d.Tag = uint8(b[5])
	// 1 byte reserved
	d.Length = uint8(b[7])
	d.Data = b[8:]
	return nil
}

func (d *Data) String() string {
	return fmt.Sprintf("Data(Enterprise: %v, Type: %#x, Tag: %#x, Length: %v, Data: % #x)",
		d.Enterprise, d.Type, d.Tag, d.Length, d.Data)
}
