package constbn

import (
	"testing"
)

func TestMuladdSmall(t *testing.T) {
	left := []Base{0}
	muladdSmall(left, zero, []Base{0})
	if left[0] != zero {
		t.Errorf("empty multiplicand was modified into %x", left)
	}

	x := []Base{Base(32), one, one}
	m := []Base{Base(32), one, one}
	muladdSmall(x, Base(57), m)

	if x[0] != Base(0x20) || x[1] != Base(0x39) || x[2] != Base(0x01) {
		t.Errorf("expected [0x20 0x39 0x01] but got %x", x)
	}
}
