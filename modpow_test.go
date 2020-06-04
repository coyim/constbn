package constbn

import (
	"reflect"
	"testing"
)

type modpowTestInstance struct {
	x      []base
	y      []byte
	blen   int
	m      []base
	ninv   base
	result []base
}

func TestModpow(t *testing.T) {
	for i, test := range modpowTestInstances[20:] {
		ourx := make([]base, len(test.x))
		copy(ourx, test.x)

		modpow(ourx, test.y[:test.blen], test.m, test.ninv)

		if !reflect.DeepEqual(ourx, test.result) {
			t.Errorf("#%d: got %x want %x", i, ourx, test.result)
			return
		}

		result := simpleModpow(test.x, test.y[:test.blen], test.m)
		if !reflect.DeepEqual(result, test.result) {
			t.Errorf("#%d: got %x want %x", i, result, test.result)
			return
		}
	}
}
