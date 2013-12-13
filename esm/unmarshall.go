package esm

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
)

func Unmarshal(data []byte, e *ESM) error {
	return nil
}

type Decoder struct {
	src io.Reader
	buf []byte
	e   *ESM
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{src: r}
}

func (d *Decoder) Buffered() []byte {
	return d.buf
}

func (d *Decoder) Decode(e *ESM) error {
	attrFields := FieldMaps(e)
	d.e = e
	err := d.parseHeader()
	if err != nil {
		return err
	}

	for {
		err = d.parseGroup(attrFields)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func FieldMaps(e *ESM) map[string]int {
	m := make(map[string]int)
	rt := reflect.TypeOf(e).Elem()
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		ft := f.Tag.Get("esm")
		st := f.Tag.Get("grouped")
		if ft != "" && st == "" {
			m[ft] = i
		}
	}
	return m
}

func (d *Decoder) parseHeader() error {
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
	n, err := d.src.Read(prefix)
	if err != nil {
		return err
	}
	if n != 24 {
		return fmt.Errorf("Could not read sufficient data for header")
	}

	// Check for format magic string
	if string(prefix[0:4]) != "TES4" {
		return fmt.Errorf("Could not find magic string for esm file")
	}

	// Think this is magic version number
	if prefix[8] != 0x01 {
		return fmt.Errorf("File version? does not match")
	}

	// At byte 20 we get a notice as to whether this file has a master file
	if prefix[20] == 0x02 {
		d.e.IsMasterFile = true
	}
	if prefix[20] == 0x0f {
		d.e.IsMasterFile = false
	}

	// the four bytes after TES4 tells us how many bytes the rest of the hedr is
	// so we'll go ahead and grab that
	hl := Decode4LE(prefix[4:8])
	buf := make([]byte, hl)
	n, err = d.src.Read(buf)
	if n != hl {
		return fmt.Errorf("HEDR record could not be retrieved in full")
	}
	if err != nil {
		return err
	}
	// turn the head content into an input
	headInput := bytes.NewBuffer(buf)

	hb := make([]byte, 6)
	n, err = headInput.Read(hb)
	if n != 6 {
		return fmt.Errorf("HEDR record and length not able to be read")
	}
	if err != nil {
		return err
	}

	hl = Decode2LE(hb[4:6])
	hb = make([]byte, hl)
	n, err = headInput.Read(hb)
	if n != hl {
		return fmt.Errorf("Could not read HEDR content")
	}
	if err != nil {
		return err
	}

	// read cnam record, has name of author
	cb := make([]byte, 6)
	n, err = headInput.Read(cb)
	if n != 6 {
		return fmt.Errorf("Could not read CNAM content")
	}
	if err != nil {
		return err
	}
	s, err := ReadString(headInput, cb[4:6])
	if err != nil {
		return err
	}

	d.e.Author = s

	// master files have no more fields at this point
	if d.e.IsMasterFile {
		return nil
	}

	// read some more items, saving those that we care about
	items, _ := ReadItems(headInput)
	if i, ok := items["MAST"]; ok {
		d.e.MasterFile = string(i[:len(i)-1])
		delete(items, "MAST")
	}
	for k, v := range items {
		d.e.Extra = append(d.e.Extra, Attribute{k, v})
	}

	return nil
}

func (d *Decoder) parseGroup(m map[string]int) error {
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
	gr := make([]byte, 24)
	n, err := d.src.Read(gr)
	if n < 24 || err == io.EOF {
		return io.EOF
	}
	if string(gr[0:4]) != "GRUP" {
		return fmt.Errorf("GRUP signifier not found")
	}
	GroupName := string(gr[8:12])
	//GroupType := RecordTypes[GroupName]

	gl := Decode4LE(gr[4:8])
	gb := make([]byte, gl-24)
	n, err = d.src.Read(gb)
	if n < gl-24 {
		return fmt.Errorf("Incomplete group section")
	}
	if err != nil && err != io.EOF {
		return err
	}
	if err == io.EOF {
		return nil
	}

	if index, ok := m[GroupName]; ok {
		fmt.Println("Supported group name", GroupName)
		fmt.Println(index)
	} else {
		fmt.Println("Unsupported group name", GroupName)
	}
	return nil
}

// Decode a 4 byte little endian integer
func Decode4LE(b []byte) int {
	return int(b[3])<<24 + int(b[2])<<16 + int(b[1])<<8 + int(b[0])
}

// Decode a 2 byte little endian integer
func Decode2LE(b []byte) int {
	return int(b[1])<<8 + int(b[0])
}

func ReadStringBuffer(buf []byte, length []byte) (string, error) {
	return ReadString(bytes.NewBuffer(buf), length)
}

// Take a length (2-4 bytes) and a  reader then reads a null
// terminated string from the reader according to the length
// passed
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
