package esm

import (
	"fmt"
)

var Fakes []*FakeDecoder

func NewFakeDecoder(name string) Decoder {
	fd := &FakeDecoder{
		Name:       name,
		Attributes: make(map[string]bool),
	}
	Fakes = append(Fakes, fd)
	return fd.Decode
}

type FakeDecoder struct {
	Name       string
	Attributes map[string]bool
}

type FakeRecord struct {
	RecordType
	RecordName string
	FormID     []byte
	EditorID   string
	Unknown    []byte
	VaryByte   byte
	Attrs      map[string][]byte
	SubRecord  []byte
}

func (fr FakeRecord) String() string {
	return fmt.Sprintf("%s: %s (%v)", fr.RecordName, fr.EditorID, fr.FormID)
}

func (fd *FakeDecoder) Decode(buf []byte) Record {
	fr := &FakeRecord{
		RecordType: UNKNOWN,
		RecordName: string(buf[0:4]),
		FormID:     buf[12:16],
		Unknown:    buf[16:20],
		VaryByte:   buf[22],
		Attrs:      make(map[string][]byte),
	}

	buf2 := buf[24:]
	// Cells start with two GRUP's before the actual CELL record
	// That's going to be a hard decoder to write
	if fr.RecordName == "GRUP" {
		fr.RecordName = "CELL"
		return fr
	}
	offset := 0
	for len(buf2) > 0 {
		// did someone fill in a record with 0's
		// found this on a WRLD record in DeadMoney.esm
		if Decode4LE(buf2[0:4]) == 0 {
			return fr
		}
		l := Decode2LE(buf2[4:6])
		offset += l

		fr.Attrs[string(buf2[0:4])] = buf2[6 : 6+l]

		if string(buf2[0:4]) == "EDID" {
			fr.EditorID = string(buf2[6 : 6+l-1])
		} else if _, ok := fd.Attributes[string(buf2[0:4])]; !ok {
			fd.Attributes[string(buf2[0:4])] = true
		}
		buf2 = buf2[6+l:]
	}

	return fr
}

func (fr *FakeRecord) ParseSubrecords(b []byte) {
	fr.SubRecord = b
}
