package esm

import (
	"fmt"
)

/*
Documentation about character classes
*/

func CLASDecode(buf []byte) Record {
	c := &Class{RecordType: CHARACTER_CLASS}
	c.FormID = buf[12:16]
	c.Unknown = buf[16:20]
	c.VaryByte = buf[22]

	buf2 := buf[24:]
	for len(buf2) > 0 {
		l := Decode2LE(buf2[4:6])
		switch string(buf2[0:4]) {
		case "EDID":
			c.EditorID = string(buf2[6 : 6+l-1])
		case "FULL":
			c.Full = string(buf2[6 : 6+l-1])
		case "DESC":
			c.Description = string(buf2[6 : 6+l-1])
		case "DATA":
			c.Data = buf2[6 : 6+l]
		case "ATTR":
			c.Attr = buf2[6 : 6+l]

		default:
			fmt.Println("mysterious thing", buf2[0:4])
		}
		buf2 = buf2[6+l:]
	}

	return c
}

type Class struct {
	RecordType
	FormID      []byte
	Unknown     []byte
	VaryByte    byte
	EditorID    string
	Full        string
	Description string
	Data        []byte
	Attr        []byte
}

func (c Class) String() string {
	return fmt.Sprintf("Character Class: %s", c.EditorID)
}
func init() {
	RecordTypes["CLAS"] = CHARACTER_CLASS
	RecordDecoders["CLAS"] = CLASDecode
}
