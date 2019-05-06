package asf

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	table := []struct {
		marshaled []byte
		want      *PresencePong
	}{
		{[]byte{0, 0, 0x8, 0xAE, 0, 0, 0, 0, 0x81, 0, 0, 0, 0, 0, 0, 0},
			&PresencePong{
				Enterprise: 2222,
				Entities:   0x81,
			}},
		{[]byte{1, 2, 3, 4, 0, 0, 0, 0, 1, 1 << 7, 0, 0, 0, 0, 0, 0},
			&PresencePong{
				Enterprise:   16909060,
				Entities:     1,
				Interactions: 1 << 7,
			}},
	}
	p := &PresencePong{}
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

func TestSupportsIPMI(t *testing.T) {
	table := []struct {
		entities uint8
		want     bool
	}{
		{0, false},
		{1 << 7, true},
		{1, false},
		{0x81, true},
	}
	for _, row := range table {
		p := &PresencePong{
			Entities: row.entities,
		}
		got := p.SupportsIPMI()
		if got != row.want {
			t.Errorf("SupportsIPMI(%#x) = %v, want %v", row.entities, got, row.want)
		}
	}
}

func TestSupportsASFV1(t *testing.T) {
	table := []struct {
		entities uint8
		want     bool
	}{
		{0, false},
		{1 << 7, false},
		{1, true},
		{0x81, true},
	}
	for _, row := range table {
		p := &PresencePong{
			Entities: row.entities,
		}
		got := p.SupportsASFV1()
		if got != row.want {
			t.Errorf("SupportsASFV1(%#x) = %v, want %v", row.entities, got, row.want)
		}
	}
}

func TestSupportsSecurityExtensions(t *testing.T) {
	table := []struct {
		interactions uint8
		want         bool
	}{
		{0, false},
		{1 << 7, true},
		{1, false},
	}
	for _, row := range table {
		p := &PresencePong{
			Interactions: row.interactions,
		}
		got := p.SupportsSecurityExtensions()
		if got != row.want {
			t.Errorf("SupportsSecurityExtensions(%#x) = %v, want %v", row.interactions, got, row.want)
		}
	}
}
