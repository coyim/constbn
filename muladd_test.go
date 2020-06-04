package constbn

import (
	"testing"
)

func TestMuladdSmall(t *testing.T) {
	left := []base{0}
	muladdSmall(left, zero, []base{0})
	if left[0] != zero {
		t.Errorf("empty multiplicand was modified into %x", left)
	}

	x := []base{base(32), one, one}
	m := []base{base(32), one, one}
	muladdSmall(x, base(57), m)

	if x[0] != base(0x20) || x[1] != base(0x39) || x[2] != base(0x01) {
		t.Errorf("expected [0x20 0x39 0x01] but got %x", x)
	}
}
