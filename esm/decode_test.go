package esm

import (
	. "github.com/acsellers/assert"
	"os"
	"testing"
)

func TestHeader(t *testing.T) {
	Within(t, func(test *Test) {
		f, e := os.Open("header_test.esm")
		test.NoError(e)
		ef, e := DecodeESM(f)
		test.NoError(e)
		test.AreEqual(ef.Header.Author, "jfader")
	})
}
