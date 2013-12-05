package esm

import (
	"fmt"
)

type MenuIcon struct {
	RecordType
	FormID   []byte
	Unknown  []byte
	VaryByte byte
	EditorID string
	Icon     string
}

func (mi MenuIcon) String() string {
	return fmt.Sprintf("MenuIcon: %s", mi.EditorID)
}

func MICNDecode(buf []byte) Record {
	mi := &MenuIcon{RecordType: MENU_ICON}
	mi.FormID = buf[12:16]
	mi.Unknown = buf[16:20]
	mi.VaryByte = buf[22]

	buf2 := buf[24:]
	for len(buf2) > 0 {
		l := Decode2LE(buf2[4:6])
		switch string(buf2[0:4]) {
		case "EDID":
			mi.EditorID = string(buf2[6 : 6+l-1])
		case "ICON":
			mi.Icon = string(buf2[6 : 6+l-1])
		default:
			fmt.Println("mysterious thing", buf2[0:4])
		}
		buf2 = buf2[6+l:]
	}

	return mi
}
func init() {
	RecordTypes["MICN"] = MENU_ICON
	RecordDecoders["MICN"] = MICNDecode
}
