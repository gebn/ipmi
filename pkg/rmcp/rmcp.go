package rmcp

const (
	// Version1 identifies RMCP v1.0 in the Version header field.
	// Lower values are considered legacy, and higher values are
	// reserved by the specification.
	Version1 uint8 = 0x06

	// Normal indicates a non-acknowledgement or "normal" message.
	Normal uint8 = 0x00

	// Ack indicates this message is acknowledging a received
	// normal message.
	Ack uint8 = 1 << 7

	// ASF identifies an RMCP message as containing an ASF-RMCP
	// payload. See the asf package within this one for details.
	ASF uint8 = 0x06

	// IPMI identifies an RMCP message as containing an IPMI
	// payload.
	IPMI uint8 = 0x07

	// OEM identifies an RMCP message as containing an OEM-defined
	// payload.
	OEM uint8 = 0x08
)
