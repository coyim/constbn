package constbn

// Int represents an arbitrarily sized integer
type Int struct {
	v []base
}

// SetBytes interprets buf as the bytes of a big-endian unsigned
// integer, sets z to that value, and returns z.
func (i *Int) SetBytes(b []byte) *Int {
	i.v = simpleDecode(b)
	return i
}

// Bytes returns the absolute value of x as a big-endian byte slice.
func (i *Int) Bytes() []byte {
	return simpleEncode(i.v)
}

// conversion to and from *big.Int
// modexp operation
