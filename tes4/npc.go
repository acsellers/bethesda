package esm

import (
	"fmt"
)

/*
NPC_ records aren't like your normal records.
Read 18 bytes, Convert bytes 4:8 to a length, and pull that
many bytes. Then decode into something useful.

The something useful isn't known yet.
*/

func NPCDecode(buf []byte) Record {
	h := make([]byte, 18)
	d := make([]byte, Decode4LE(buf[4:8]))
	n := &NPC{RecordType: NPC_}
	n.Header = h[8:]
	n.Data = d
	return n
}

type NPC struct {
	RecordType
	Header []byte
	Data   []byte
}

func (n NPC) String() string {
	return fmt.Sprintf("NPC")
}

func init() {
	RecordTypes["NPC_"] = NPC_
	RecordDecoders["NPC_"] = NPCDecode
}
