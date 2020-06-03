package constbn

func byteLen(a []base) base {
	return baseLen(a) << 2
}

func baseLen(a []base) base {
	return (a[0] + 31) >> 5
}

/*
 * Convert a modular integer to Montgomery representation. The integer x[]
 * MUST be lower than m[], but with the same announced bit length.
 */
func toMonty(x, m []base) {
	for k := baseLen(m); k > zero; k-- {
		muladdSmall(x, zero, m)
	}
}
