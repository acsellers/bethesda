package esm

import (
	"fmt"
)

/*

*/

func SOUNDecode(buf []byte) Record {
	c := &Sound{RecordType: SOUND}
	c.FormID = buf[12:16]
	c.Unknown = buf[16:20]
	c.VaryByte = buf[22]

	buf2 := buf[24:]
	for len(buf2) > 0 {
		l := Decode2LE(buf2[4:6])
		switch string(buf2[0:4]) {
		case "EDID":
			c.EditorID = string(buf2[6 : 6+l-1])
		case "FNAM":
			c.FileName = string(buf2[6 : 6+l-1])
		case "RNAM":
			c.RName = string(buf2[6 : 6+l-1])

		case "OBND":
			c.OBND = buf2[6 : 6+l]
		case "SNDD":
			c.SoundData = buf2[6 : 6+l]

		default:
			fmt.Println("mysterious thing", buf2[0:4])
		}
		buf2 = buf2[6+l:]
	}

	return c
}

type Sound struct {
	RecordType
	FormID    []byte
	Unknown   []byte
	VaryByte  byte
	EditorID  string
	OBND      []byte
	FileName  string
	RName     string
	SoundData []byte
}

func (s Sound) String() string {
	return fmt.Sprintf("Sound: %s", s.EditorID)
}
func init() {
	RecordTypes["SOUN"] = SOUND
	RecordDecoders["SOUN"] = SOUNDecode
}
