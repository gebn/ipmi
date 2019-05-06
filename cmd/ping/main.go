// Package main implements a simple RMCP Presence Ping tool.
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/gebn/ipmi/pkg/rmcp"
	"github.com/gebn/ipmi/pkg/rmcp/asf"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %v <bmc address>\n", os.Args[0])
		return
	}

	// connect
	raddr, err := net.ResolveUDPAddr("udp", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	log.Printf("Connected to %v on %v", conn.RemoteAddr(), conn.LocalAddr())

	// write
	ping := &asf.Data{
		Enterprise: asf.Enterprise,
		Type:       asf.TypePresencePing,
	}
	request := &rmcp.Packet{
		Version:  rmcp.Version1,
		Sequence: 5,
		Class:    rmcp.Normal | rmcp.ASF,
		Data:     ping.Marshal(),
	}

	if err := conn.SetWriteDeadline(time.Now().Add(time.Second * 2)); err != nil {
		log.Fatal(err)
	}
	marshaled := request.Marshal()
	n, err := conn.Write(marshaled)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %v bytes", n)
	if n != len(marshaled) {
		log.Fatal("Wrote incomplete message")
	}

	// read
	if err := conn.SetReadDeadline(time.Now().Add(time.Second * 5)); err != nil {
		log.Fatal(err)
	}
	buffer := make([]byte, 50)
	n, _, err = conn.ReadFromUDP(buffer)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Read %v bytes", n)

	// parse
	response := &rmcp.Packet{}
	if err := response.Parse(buffer[:n]); err != nil {
		log.Fatal(err)
	}
	responseData := &asf.Data{}
	if err := responseData.Parse(response.Data); err != nil {
		log.Fatal(err)
	}
	pong := &asf.PresencePong{}
	if err := pong.Parse(responseData.Data); err != nil {
		log.Fatal(err)
	}

	log.Printf("Supports IPMI: %v", pong.SupportsIPMI())
	log.Printf("Supports ASF V1.0: %v", pong.SupportsASFV1())
	log.Printf("Supports security extensions: %v", pong.SupportsSecurityExtensions())
}
