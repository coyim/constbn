package constbn

/*
 * Compute a modular exponentiation. x[] MUST be an integer modulo m[]
 * (same announced bit length, lower value). m[] MUST be odd. The
 * exponent is in big-endian unsigned notation, over 'elen' bytes. The
 * "m0i" parameter is equal to -(1/m0) mod 2^31, where m0 is the least
 * significant value word of m[] (this works only if m[] is an odd
 * integer).
 */

func modpow(x []base, e []byte, m []base, m0i base) {
	elen := len(e)
	t1 := make([]base, len(m))
	t2 := make([]base, len(m))

	mlen := baseLen(m)

	copy(t1, x[:mlen])
	toMonty(t1, m)
	zeroize(x, m[0])
	x[1] = one
	for k := zero; k < base(elen<<3); k++ {
		ctl := base((e[elen-1-int(k>>3)] >> (k & 7))) & 1
		montmul(t2, x, t1, m, m0i)
		ccopy(ctl, x, t2, mlen)
		montmul(t2, t1, t1, m, m0i)
		copy(t1, t2[:mlen])
	}
}
