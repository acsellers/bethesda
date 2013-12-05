package esm

import (
	"fmt"
)

/*
Game settings start with the identifier GMST, and can be ints, floats, or strings
This is indicated by the first letter of the EditorID.

Structure of a Game Setting

4 bytes (GMST)
4 byte length (from EDID to end)
4 zero bytes
4 bytes for FormID
4 bytes unknown (some sort of global object id?)
2 byte (0F 00)
2 bytes (00 00, 01 00, 02 00, 03 00)
4 bytes (EDID)
2 byte length
? bytes EditorID (string)
4 bytes (DATA)
2 byte length
? bytes Data (int, float32, string)
*/
type GameSetting struct {
	RecordType
	FormID   []byte
	Unknown  []byte
	VaryByte byte
	EditorID string
	Float    float32
	Int      int
	Str      string
	RawData  []byte
}

func GMSTDecode(buf []byte) Record {
	gm := &GameSetting{RecordType: GAME_SETTING}
	gm.FormID = buf[12:16]
	gm.Unknown = buf[16:20]
	gm.VaryByte = buf[22]
	buf2 := buf[24:]
	for len(buf2) > 0 {
		l := Decode2LE(buf2[4:6])
		switch string(buf2[0:4]) {
		case "EDID":
			gm.EditorID = string(buf2[6 : 6+l-1])
		case "DATA":
			gm.RawData = buf2[6 : 6+l]
		default:
			fmt.Println("mysterious thing", buf2[0:4])
		}
		buf2 = buf2[6+l:]
	}
	return gm
}

func (gm GameSetting) String() string {
	return fmt.Sprintf("Setting: %s", gm.EditorID)
}

func init() {
	RecordTypes["GMST"] = GAME_SETTING
	RecordDecoders["GMST"] = GMSTDecode
}
