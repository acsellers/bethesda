package esm

import (
	"bytes"
	"fmt"
	"io"
)

type Decoder func(buffer []byte) Record

type Record interface {
	Type() RecordType
	String() string
	ParseSubrecords([]byte)
}

type RecordType int

func (t RecordType) Type() RecordType {
	return t
}
func (t RecordType) ParseSubrecords(b []byte) {
	panic("NONONO")
}

const (
	HEADER RecordType = iota
	GROUP
	GAME_SETTING
	TEXTURE_SET
	MENU_ICON
	NPC_
	GLOBAL
	CHARACTER_CLASS
	FACTION
	RACE
	SOUND
	MAGIC_EFFECT
	ENCHANTMENT
	UNKNOWN
)

var RecordDecoders = make(map[string]Decoder)
var ComplexDecoding = map[string]bool{
//"DIAL": true, // NPC_ can have embedded GRUP records, not ready for this yet
//"CELL": true, // Skip for simplicity
//"WRLD": true, // Ditto
}
var RecordTypes = make(map[string]RecordType)

// Decode a 4 byte length in the TES4 Length format
func Decode4LE(b []byte) int {
	return int(b[3])<<24 + int(b[2])<<16 + int(b[1])<<8 + int(b[0])
}

// Decode a 2 byte length in the TES4 Length format
func Decode2LE(b []byte) int {
	return int(b[1])<<8 + int(b[0])
}
func ReadStringBuffer(buf []byte, length []byte) (string, error) {
	return ReadString(bytes.NewBuffer(buf), length)
}
func ReadString(input io.Reader, length []byte) (string, error) {
	var sl int
	if len(length) == 2 {
		sl = Decode2LE(length)
	} else if len(length) == 4 {
		sl = Decode4LE(length)
	} else {
		return "", fmt.Errorf("Length is not recognized, should be 2 or 4 bytes")
	}

	sb := make([]byte, sl)
	n, err := input.Read(sb)
	if n != sl {
		return "", fmt.Errorf("Could not read sufficient bytes for a string")
	}
	if err != nil {
		return "", err
	}

	// ESM strings are null terminated, we don't need the null
	return string(sb[:len(sb)-1]), nil
}

func ReadItems(input io.Reader) (map[string][]byte, error) {
	items := make(map[string][]byte)
	for {
		ip := make([]byte, 6)
		n, err := input.Read(ip)
		if err == io.EOF || n < 6 {
			break
		}
		if err != nil {
			return items, err
		}
		il := Decode2LE(ip[4:6])
		id := make([]byte, il)
		n, err = input.Read(id)

		items[string(ip[:4])] = id
		if il > n {
			return items, fmt.Errorf("Item underrun")
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return items, err
		}
	}
	return items, nil
}

type HeaderRecord struct {
	RecordType
	IsMasterFile bool
	MasterFile   string
	RawRecord    []byte
	FileVersion  int
	Author       string
	OtherItems   map[string][]byte
}

type GroupRecord struct {
	RecordType
	GroupName string
	GroupType RecordType
	Records   []Record
}

func DecodeHeader(input io.Reader) (*HeaderRecord, error) {
	/*
				Structure of a header

				4 bytes (TES4)
					4 bytes, length from HEDR to end of HEDR
					0x01
		      3 zero bytes
					8 zero bytes
					0x02 for master files, 0x0F for non-masters
					3 zero bytes
				4 bytes (HEDR)
					2 byte length (0C 00 usually)
					12 unknown bytes
				4 bytes (CNAM)
					2 byte length
					String (null-terminated) jfader for NV dlc, ipeley for NV
				- for non-master files
					4 bytes (MAST)
					  2 byte length
					  String of filename for master file, case sensitive
					4 bytes (DATA)
					  2 byte length
					  8 zero bytes
					4 bytes (ONAM)
					  2 byte length
					  variable data
	*/

	// first part of the file, starts with TES4, prior to HEDR record
	prefix := make([]byte, 24)
	n, err := input.Read(prefix)
	if err != nil {
		return nil, err
	}
	if n != 24 {
		return nil, fmt.Errorf("Could not read sufficient data for header")
	}

	// Check for format magic string
	if string(prefix[0:4]) != "TES4" {
		return nil, fmt.Errorf("Could not find magic string for esm file")
	}

	// Think this is magic version number
	if prefix[8] != 0x01 {
		return nil, fmt.Errorf("File version? does not match")
	}

	// At byte 20 we get a notice as to whether this file has a master file
	hr := &HeaderRecord{RecordType: HEADER}
	if prefix[20] == 0x02 {
		hr.IsMasterFile = true
	}
	if prefix[20] == 0x0f {
		hr.IsMasterFile = false
	}

	// the four bytes after TES4 tells us how many bytes the rest of the hedr is
	// so we'll go ahead and grab that
	hl := Decode4LE(prefix[4:8])
	buf := make([]byte, hl)
	n, err = input.Read(buf)
	if n != hl {
		return nil, fmt.Errorf("HEDR record could not be retrieved in full")
	}
	if err != nil {
		return nil, err
	}
	// turn the head content into an input
	headInput := bytes.NewBuffer(buf)

	hb := make([]byte, 6)
	n, err = headInput.Read(hb)
	if n != 6 {
		return nil, fmt.Errorf("HEDR record and length not able to be read")
	}
	if err != nil {
		return nil, err
	}

	hl = Decode2LE(hb[4:6])
	hb = make([]byte, hl)
	n, err = headInput.Read(hb)
	if n != hl {
		return nil, fmt.Errorf("Could not read HEDR content")
	}
	if err != nil {
		return nil, err
	}
	hr.RawRecord = hb

	// read cnam record, has name of author
	cb := make([]byte, 6)
	n, err = headInput.Read(cb)
	if n != 6 {
		return nil, fmt.Errorf("Could not read CNAM content")
	}
	if err != nil {
		return nil, err
	}
	s, err := ReadString(headInput, cb[4:6])
	if err != nil {
		return nil, err
	}

	hr.Author = s

	// master files have no more fields at this point
	if hr.IsMasterFile {
		return hr, nil
	}

	// read some more items, saving those that we care about
	items, _ := ReadItems(headInput)
	if i, ok := items["MAST"]; ok {
		hr.MasterFile = string(i[:len(i)-1])
		delete(items, "MAST")
	}
	hr.OtherItems = items

	return hr, nil
}

func DecodeGroup(input io.Reader) (GroupRecord, error) {
	/*
	  Structure of a Group Record
	  4 bytes (GRUP)
	    4 byte length (LE)
	    4 bytes, type name
	    4 zero bytes
	    4 non-zero bytes
	    4 bytes 01000000 on master, zero otherwise
	  a number of child records
	*/
	g := GroupRecord{RecordType: GROUP}
	gr := make([]byte, 24)
	n, err := input.Read(gr)
	if n < 24 || err == io.EOF {
		return g, io.EOF
	}
	if string(gr[0:4]) != "GRUP" {
		return g, fmt.Errorf("GRUP signifier not found")
	}
	g.GroupName = string(gr[8:12])
	g.GroupType = RecordTypes[g.GroupName]

	gl := Decode4LE(gr[4:8])
	gb := make([]byte, gl-24)
	n, err = input.Read(gb)
	if n < gl-24 {
		return g, fmt.Errorf("Incomplete group section")
	}
	if err != nil && err != io.EOF {
		return g, err
	}
	if d, ok := RecordDecoders[g.GroupName]; ok {
		for len(gb) > 0 {
			rl := Decode4LE(gb[4:8]) + 24
			if string(gb[0:4]) == "GRUP" {
				rl -= 24
			}
			r := d(gb[:rl])
			gb = gb[rl:]

			g.Records = append(g.Records, r)
		}
		//} else if _, ok := SkipDecoding[g.GroupName]; ok {
		//	fmt.Println("Skipping", g.GroupName, "Section")
	} else {
		fmt.Println("Loading fake decoder for", g.GroupName)
		d := NewFakeDecoder(g.GroupName)
		RecordDecoders[g.GroupName] = d

		for len(gb) > 0 {
			rl := Decode4LE(gb[4:8]) + 24
			if string(gb[0:4]) == "GRUP" {
				rl -= 24
			}

			fmt.Println(rl, len(gb))
			r := d(gb[:rl])
			gb = gb[rl:]
			g.Records = append(g.Records, r)
		}
	}

	return g, nil
}

type ESMFile struct {
	Header *HeaderRecord
	Groups []GroupRecord
}

func DecodeESM(input io.Reader) (*ESMFile, error) {
	ef := &ESMFile{}
	h, err := DecodeHeader(input)
	if err != nil {
		return nil, err
	}
	ef.Header = h

	group, err := DecodeGroup(input)
	for err == nil {
		ef.Groups = append(ef.Groups, group)
		group, err = DecodeGroup(input)
	}

	if err != io.EOF {
		return ef, err
	}

	return ef, nil
}
