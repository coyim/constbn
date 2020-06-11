package constbn

import (
	"reflect"
	"testing"
)

func TestModpowOpt(t *testing.T) {
	// Testing too small buffer
	tt := modpowTestInstances[0]
	ourx := make([]Base, len(tt.x))
	copy(ourx, tt.x)
	smalltmp := make([]Base, 1)

	res := ModpowOpt(ourx, tt.y[:tt.blen], tt.m, tt.ninv, smalltmp)
	if res != zero {
		t.Errorf("expected small tmp buffer to fail")
	}

	for i, test := range modpowTestInstances {
		ourx := make([]Base, len(test.x))
		copy(ourx, test.x)
		tmp1 := make([]Base, 3*baseLenWithHeader(test.m))
		_ = ModpowOpt(ourx, test.y[:test.blen], test.m, test.ninv, tmp1)
		if !reflect.DeepEqual(ourx, test.result) {
			t.Errorf("#%d: got %x want %x", i, ourx, test.result)
			return
		}

		ourx2 := make([]Base, len(test.x))
		copy(ourx2, test.x)
		tmp2 := make([]Base, 4*baseLenWithHeader(test.m))
		_ = ModpowOpt(ourx2, test.y[:test.blen], test.m, test.ninv, tmp2)
		if !reflect.DeepEqual(ourx2, test.result) {
			t.Errorf("#%d: got %x want %x", i, ourx2, test.result)
			return
		}

		ourx3 := make([]Base, len(test.x))
		copy(ourx3, test.x)
		tmp3 := make([]Base, 50*baseLenWithHeader(test.m))
		_ = ModpowOpt(ourx3, test.y[:test.blen], test.m, test.ninv, tmp3)
		if !reflect.DeepEqual(ourx3, test.result) {
			t.Errorf("#%d: got %x want %x", i, ourx3, test.result)
			return
		}

		result := simpleModpowOpt(test.x, test.y[:test.blen], test.m)
		if !reflect.DeepEqual(result, test.result) {
			t.Errorf("#%d: got %x want %x", i, result, test.result)
			return
		}
	}
}
