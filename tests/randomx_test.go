package tests

import (
	"encoding/hex"
	"strconv"
	"testing"

	"github.com/gcarreno/go-randomx"
)

type TestVector struct {
	Input  []byte
	Nonce  int64
	Output string // hex-encoded expected hash
}

var testVectors = []TestVector{
	{
		Input:  []byte("hello"),
		Nonce:  12345,
		Output: "1380b8342891a69bc27e8be7f8845061fe3588a94b1da4750cc525be686f03d3", // this would be from a known output
	},
	// Add more
}

func TestRandomX(t *testing.T) {
	for _, tv := range testVectors {
		nonce := []byte(strconv.FormatInt(tv.Nonce, 10))
		rx, err := randomx.NewRandomX(nonce)
		if err != nil {
			t.Errorf("Error getting NewRandomX: %v", err)
		}

		got, err := rx.ComputeHash(tv.Input)
		if err != nil {
			t.Errorf("Error computing hash: %v", err)
		}

		if hex.EncodeToString(got) != tv.Output {
			t.Errorf("Input: %q Nonce: %d\nExpected: %s\nGot:      %x",
				tv.Input,
				tv.Nonce,
				tv.Output,
				got,
			)
		}
	}
}
