package constbn

import (
	"testing"
)

func TestAPISet(t *testing.T) {
	v := new(Int).SetBytes([]byte{0x01})
	v2 := new(Int).SetBytes([]byte{0x02})

	v.Set(v2)

	if v.GetBigInt().String() != "2" {
		t.Errorf("Int.Set() doesn't work - got: %#x", v.GetBigInt().String())
	}
}
