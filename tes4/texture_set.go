package esm

import (
	"fmt"
)

type TextureSet struct {
	RecordType
	FormID    []byte
	Unknown   []byte
	VaryByte  byte
	EditorID  string
	DecalData []byte
	Textures  [6]string
}

func (ts TextureSet) String() string {
	return fmt.Sprintf("TextureSet: %s", ts.EditorID)
}

func TXSTDecode(buf []byte) Record {
	ts := &TextureSet{RecordType: TEXTURE_SET}
	ts.FormID = buf[12:16]
	ts.Unknown = buf[16:20]
	ts.VaryByte = buf[22]

	buf2 := buf[24:]
	for len(buf2) > 0 {
		l := Decode2LE(buf2[4:6])
		switch string(buf2[0:4]) {
		case "EDID":
			ts.EditorID = string(buf2[6 : 6+l-1])
		case "OBND":
			// I would like to know what the purpose of this is
			// I found it included in the files, but with a length of 12 and full of zero bytes
			// In that case, we'll just ignore it
		case "DNAM":
			// Same thing as OBND, except it has a length of 2
		case "DODT":
			// Decal Object Data, for now we'll just pretend it's magic bytes
			// At some point, we can reverse engineer the format and put it into a DecalData object
			ts.DecalData = buf2[6 : 6+l]
		case "TX00":
			ts.Textures[0] = string(buf2[6 : 6+l-1])
		case "TX01":
			ts.Textures[1] = string(buf2[6 : 6+l-1])
		case "TX02":
			ts.Textures[2] = string(buf2[6 : 6+l-1])
		case "TX03":
			ts.Textures[3] = string(buf2[6 : 6+l-1])
		case "TX04":
			ts.Textures[4] = string(buf2[6 : 6+l-1])
		case "TX05":
			ts.Textures[5] = string(buf2[6 : 6+l-1])
		default:
			fmt.Println("mysterious thing", buf2[0:4])
		}
		buf2 = buf2[6+l:]
	}

	return ts
}
func init() {
	RecordTypes["TXST"] = TEXTURE_SET
	RecordDecoders["TXST"] = TXSTDecode
}
