// Package asf implements components of RMCP specific to IANA Enterprise Number
// 4542 (Alerting Specifications Forum). The existence of this package *within*
// RMCP, when ASF defined RMCP, is a little idiosyncratic however, in practice,
// RMCP is tightly coupled to the ASF implementation.
//
// In theory, OEMs can plonk their own Enterprise Number in the first field of
// the payload and define their own structure, however this is yet to be seen,
// and is certainly not required for IPMI compliance.
package asf

const (
	// Enterprise is the IANA-assigned Enterprise Number of the ASF-RMCP
	// organisation.
	Enterprise uint32 = 4542

	// TypePresencePong is the message type of the response to a Presence Ping
	// message. It indicates the sender is ASF-RMCP-aware.
	TypePresencePong uint8 = 0x40

	// TypePresencePing is a message type sent to a managed client to solicit a
	// Presence Pong response. Clients may ignore this is the RMCP version is
	// unsupported. Sending this message with a sequence number <255 is the
	// recommended way of finding out whether an implementation sends RMCP ACKs.
	// (Super Micro does not).
	//
	// Systems implementing IPMI must respond to this ping to conform to the
	// spec, so it is a good, lightweight substitute for an ICMP ping.
	TypePresencePing uint8 = 0x80
)
