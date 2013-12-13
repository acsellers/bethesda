package esm

import (
	"bytes"
	. "github.com/acsellers/assert"
	"io"
	"os"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	Within(t, func(test *Test) {
		d := NewDecoder(&bytes.Buffer{})
		test.AreEqual(d.Decode(&ESM{}), io.EOF)

		f, e := os.Open("header_test.esm")
		test.NoError(e)
		d = NewDecoder(f)
		var tesm ESM
		e = d.Decode(&tesm)
		test.NoError(e)
		test.AreEqual(tesm.Author, "jfader")
		test.IsFalse(tesm.IsMasterFile)
		test.AreEqual(tesm.MasterFile, "FalloutNV.esm")
	})
}
