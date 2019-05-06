package asf

import (
	"bytes"
	"reflect"
	"testing"
)

func TestMarshalParse(t *testing.T) {
	table := []struct {
		data      *Data
		marshaled []byte
	}{
		{&Data{
			Enterprise: 1234,
			Type:       TypePresencePing,
			Tag:        22,
			Data:       []byte{}, // nil makes test fail
		}, []byte{0, 0, 0x4, 0xD2, 0x80, 22, 0, 0}},
		{&Data{
			Enterprise: Enterprise,
			Type:       TypePresencePong,
			Tag:        25,
			Length:     2,
			Data:       []byte{1, 2},
		}, []byte{0, 0, 0x11, 0xBE, 0x40, 25, 0, 2, 1, 2}},
	}
	for _, row := range table {
		// check data is marshaled correctly
		marshaled := row.data.Marshal()
		if !bytes.Equal(marshaled, row.marshaled) {
			t.Errorf("Marshal(%v) = %v, want %v", row.data, marshaled, row.marshaled)
			continue
		}

		// and check it unmarshals to the original data
		d := &Data{}
		if err := d.Parse(marshaled); err != nil {
			t.Errorf("Parse(%v) = %v, want no error", marshaled, err)
			continue
		}
		if !reflect.DeepEqual(d, row.data) {
			t.Errorf("Parse(Marshal(%v)) = %v, want %v", row.data, d, row.data)
		}
	}
}
