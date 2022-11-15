package internal

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

/*
         zy   xwvutsrq    ponmlkji    hgfedcba
0b00000000, 0b00000000, 0b00000000, 0b00000000

*/

func Test_Pack(t *testing.T) {
	tests := []struct {
		input    string
		expected [4]byte
	}{
		{
			"a",
			[4]byte{0b00000000, 0b00000000, 0b00000000, 0b00000001},
		},
		{
			"aaaaaaa",
			[4]byte{0b00000000, 0b00000000, 0b00000000, 0b00000001},
		},
		{
			"zaza",
			[4]byte{0b00000010, 0b00000000, 0b00000000, 0b00000001},
		},
		{
			"hello world",
			[4]byte{0b00000000, 0b01000010, 0b01001000, 0b10011000},
		},
		{
			"palabras",
			[4]byte{0b00000000, 0b00000110, 0b10001000, 0b00000011},
		},
		{
			"bala",
			[4]byte{0b00000000, 0b00000110, 0b10001000, 0b00000011},
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			actual := pack(test.input)
			if 0 != bytes.Compare(actual[:], test.expected[:]) {
				t.Errorf("does not match\nexpected: %s\n  actual: %s", fmtPacked(test.expected), fmtPacked(actual))
			}
		})
	}

}

func fmtPacked(in [4]byte) string {
	buf := make([]string, 4)
	for i, b := range in {
		buf[i] = fmt.Sprintf("%08b", b)
	}
	return strings.Join(buf, " ")
}
