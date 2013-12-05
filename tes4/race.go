package esm

import (
	"fmt"
)

/*
Documentation about Races
*/

func RACEDecode(buf []byte) Record {
	c := &Race{RecordType: RACE}
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
		case "MODL":
			c.Model = string(buf2[6 : 6+l-1])
		case "ICON":
			c.Icon = string(buf2[6 : 6+l-1])

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

type Race struct {
	RecordType
	FormID      []byte
	Unknown     []byte
	VaryByte    byte
	EditorID    string
	Data        []byte
	Attr        []byte
	Description string
	Model       string
	Icon        string

	Full string
	/*
		PNAM
		FNAM
		SNAM
		VTCK
		CNAM
		INDX
		HNAM
		ENAM
		NAM2
		NAM0
		MNAM
		FGGS
		FGGA
		UNAM
		NAM1
		FGTS
	*/
}

func (c Race) String() string {
	return fmt.Sprintf("Character Class: %s", c.EditorID)
}
func init() {
	//	RecordTypes["CLAS"] = CHARACTER_CLASS
	//	RecordDecoders["CLAS"] = CLASDecode
}
