package esm

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

/*
Global variables are identified by GLOB records.

Structure of a Global Variable:

4 bytes (GLOB)
4 byte length
4 zero bytes or 40 00 00 00 (constant)
4 bytes for FormID
4 bytes unknown
4 bytes (EDID)
2 byte length
Editor ID string
4 bytes (FNAM)
2 byte length
1 byte character (s, f, l)
4 bytes (FLTV)
2 byte length
value (float, short, etc.)
*/

type Global struct {
	RecordType
	FormID   []byte
	Unknown  []byte
	Constant bool
	VaryByte byte
	EditorID string
	VarType  string
	Float    float32
	Short    int16
	Long     int32
	RawData  []byte
}

func GLOBDecode(buf []byte) Record {
	g := &Global{RecordType: GLOBAL}
	g.FormID = buf[12:16]
	g.Constant = 0x40&buf[8] != 0x0
	g.Unknown = buf[16:20]
	g.VaryByte = buf[22]
	buf2 := buf[24:]
	for len(buf2) > 0 {
		l := Decode2LE(buf2[4:6])
		switch string(buf2[0:4]) {
		case "EDID":
			g.EditorID = string(buf2[6 : 6+l-1])
		case "FNAM":
			g.VarType = string(buf2[6:7])
		case "FLTV":
			switch g.VarType {
			case "f":
				binary.Read(bytes.NewBuffer(buf2[6:10]), binary.LittleEndian, &g.Float)
			case "s":
				binary.Read(bytes.NewBuffer(buf2[6:10]), binary.LittleEndian, &g.Float)
				g.Short = int16(g.Float)
			default:
				binary.Read(bytes.NewBuffer(buf2[6:10]), binary.LittleEndian, &g.Float)
				g.Short = int16(g.Float)
			}
		default:
			fmt.Println("mysterious thing", buf2[0:4])
		}
		buf2 = buf2[6+l:]
	}
	return g

}
func (g Global) String() string {
	return fmt.Sprintf("Global Variable: %s", g.EditorID)
}

func init() {
	RecordTypes["GLOB"] = GLOBAL
	RecordDecoders["GLOB"] = GLOBDecode
}
