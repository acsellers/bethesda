package esm

import (
	"fmt"
)

/*
Yes, there is magic in fallout, or at least repurposed magic...
*/
type MagicEffect struct {
	RecordType
	FormID      []byte
	Unknown     []byte
	VaryByte    byte
	EditorID    string
	Full        string
	Description string
	Data        []byte
}

func MGEFDecode(buf []byte) Record {
	gm := &MagicEffect{RecordType: MAGIC_EFFECT}
	gm.FormID = buf[12:16]
	gm.Unknown = buf[16:20]
	gm.VaryByte = buf[22]
	buf2 := buf[24:]
	for len(buf2) > 0 {
		l := Decode2LE(buf2[4:6])
		switch string(buf2[0:4]) {
		case "EDID":
			gm.EditorID = string(buf2[6 : 6+l-1])
		case "FULL":
			gm.Full = string(buf2[6 : 6+l-1])
		case "DESC":
			gm.Description = string(buf2[6 : 6+l-1])

		case "DATA":
			gm.Data = buf2[6 : 6+l]
		default:
			fmt.Println("mysterious thing", buf2[0:4])
		}
		buf2 = buf2[6+l:]
	}
	return gm
}

func (gm MagicEffect) String() string {
	return fmt.Sprintf("MagicEffect: %s", gm.EditorID)
}

func init() {
	RecordTypes["MGEF"] = MAGIC_EFFECT
	RecordDecoders["MGEF"] = MGEFDecode
}
