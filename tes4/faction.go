package esm

import (
	"fmt"
)

/*

*/

func FACTDecode(buf []byte) Record {
	c := &Faction{RecordType: FACTION}
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
		case "DATA":
			c.Data = buf2[6 : 6+l]
		case "XNAM":
			c.XNAM = append(c.XNAM, buf2[6:6+l])

		default:
			fmt.Println("mysterious thing", buf2[0:4])
		}
		buf2 = buf2[6+l:]
	}

	return c
}

type Faction struct {
	RecordType
	FormID   []byte
	Unknown  []byte
	VaryByte byte
	EditorID string
	Full     string
	Data     []byte
	XNAM     [][]byte
}

func (f Faction) String() string {
	return fmt.Sprintf("Faction: %s", f.EditorID)
}
func init() {
	RecordTypes["FACT"] = FACTION
	RecordDecoders["FACT"] = FACTDecode
}
