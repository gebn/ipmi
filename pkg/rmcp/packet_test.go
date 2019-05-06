package rmcp

import (
	"bytes"
	"reflect"
	"testing"
)

func TestIsAck(t *testing.T) {
	table := []struct {
		class uint8
		want  bool
	}{
		{Normal, false},
		{Ack, true},
	}
	for _, row := range table {
		p := &Packet{
			Class: row.class,
		}
		got := p.IsAck()
		if got != row.want {
			t.Errorf("IsAck(%v) = %v, want %v", row.class, got, row.want)
		}
	}
}

func TestMarshal(t *testing.T) {
	table := []struct {
		packet Packet
		want   []byte
	}{
		{Packet{}, make([]byte, 4)},
		{Packet{
			Version:  Version1,
			Sequence: 22,
			Class:    Normal | IPMI,
			Data:     []byte{0x23, 0x24},
		}, []byte{Version1, 0x00, 22, Normal | IPMI, 0x23, 0x24}},
	}
	for _, row := range table {
		got := row.packet.Marshal()
		if !bytes.Equal(got, row.want) {
			t.Errorf("Marshal(%v) = %v, want %v", row.packet, got, row.want)
		}
	}
}

func TestParse(t *testing.T) {
	table := []struct {
		marshaled []byte
		want      *Packet
	}{
		{[]byte{}, nil},
		{[]byte{0x06, 0x00, 0xFF, 0x06},
			&Packet{
				Version:  Version1,
				Sequence: 255,
				Class:    Normal | ASF,
				Data:     []byte{},
			},
		},
		{[]byte{0x06, 0x00, 22, Normal | IPMI, 0x23, 0x24},
			&Packet{
				Version:  Version1,
				Sequence: 22,
				Class:    Normal | IPMI,
				Data:     []byte{0x23, 0x24},
			},
		},
	}
	p := &Packet{}
	for _, row := range table {
		err := p.Parse(row.marshaled)
		switch {
		case err != nil && row.want != nil:
			t.Errorf("Parse(%v) = %v, want nil", row.marshaled, err)
		case err == nil && !reflect.DeepEqual(p, row.want):
			t.Errorf("Parse(%v) = %v, want %v", row.marshaled, p, row.want)
		}
	}
}
