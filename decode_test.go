package constbn

import (
	"reflect"
	"testing"
)

type decodeTestInstance struct {
	m1   []Base
	bm   []byte
	blen int
}

func TestDecode(t *testing.T) {
	for i, test := range decodeTestInstances {
		ourm1 := make([]Base, 35)

		Decode(ourm1, test.bm[:test.blen])

		if !reflect.DeepEqual(ourm1, test.m1) {
			t.Errorf("#%d: got %x want %x", i, ourm1, test.m1)
		}
	}
}
