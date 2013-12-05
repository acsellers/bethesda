package esm

import (
	"fmt"
)

/*
TODO Document Enchantments
*/
type Enchantment struct {
	RecordType
	FormID   []byte
	Unknown  []byte
	VaryByte byte
	EditorID string
	Full     string
	Enit     []byte
	Efid     [][]byte
	Efit     [][]byte
}

func ENCHDecode(buf []byte) Record {
	gm := &Enchantment{RecordType: ENCHANTMENT}
	gm.FormID = buf[12:16]
	gm.Unknown = buf[16:20]
	gm.VaryByte = buf[22]
	buf2 := buf[24:]
	for len(buf2) > 0 {
		l := Decode2LE(buf2[4:6])
		switch string(buf2[0:4]) {
		case "EDID":
			gm.EditorID = string(buf2[6 : 6+l-1])
		case "Full":
			gm.EditorID = string(buf2[6 : 6+l-1])

		case "ENIT":
			gm.Enit = buf2[6 : 6+l]

		case "EFIT":
			gm.Efit = append(gm.Efit, buf2[6:6+l])
		case "EFID":
			gm.Efid = append(gm.Efid, buf2[6:6+l])

		default:
			fmt.Println("mysterious thing", buf2[0:4])
		}
		buf2 = buf2[6+l:]
	}
	return gm
}

func (gm Enchantment) String() string {
	return fmt.Sprintf("Enchantment: %s", gm.EditorID)
}

func init() {
	RecordTypes["ENCH"] = ENCHANTMENT
	RecordDecoders["ENCH"] = ENCHDecode
}
