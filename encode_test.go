package constbn

import (
	"reflect"
	"testing"
)

type encodeTestInstance struct {
	expected []byte
	input    []Base
	blen     int
}

func TestEncode(t *testing.T) {
	for i, test := range encodeTestInstances {
		ourbm := make([]byte, 128)

		Encode(ourbm[:test.blen], test.input)

		if !reflect.DeepEqual(ourbm, test.expected) {
			t.Errorf("#%d: got %x want %x", i, ourbm, test.expected)
			return
		}
	}

	for i, test := range decodeTestInstances {
		ourbm := make([]byte, 128)

		Encode(ourbm[:test.blen], test.m1)

		if !reflect.DeepEqual(ourbm, test.bm) {
			t.Errorf("#%d: got %x want %x", i, ourbm, test.bm)
			return
		}
	}
}
