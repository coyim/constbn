package main

// #include <stdlib.h>
// #include <stdint.h>
import "C"
import (
	"crypto/rand"
	"errors"
	"io"
	"math/big"
	"runtime"

	"github.com/coyim/constbn"
)

// if you change these values, change them in dut_go.c aswell!
const chunksize = 40
const measurements = 1e6

// either 0x00 or 0x01
func randombit() byte {
	b := make([]byte, 1)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b[0] & 0x01
}

var (
	x   []constbn.Base
	m   []constbn.Base
	m0i constbn.Base
	tmp []constbn.Base
)

//export init_dut
func init_dut() {
	g1 := big.NewInt(2).Bytes()
	x = make([]constbn.Base, (len(g1)/2)+2)
	constbn.Decode(x, g1)

	p, _ := new(big.Int).SetString(
		"FFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD1"+
			"29024E088A67CC74020BBEA63B139B22514A08798E3404DD"+
			"EF9519B3CD3A431B302B0A6DF25F14374FE1356D6D51C245"+
			"E485B576625E7EC6F44C42E9A637ED6B0BFF5CB6F406B7ED"+
			"EE386BFB5A899FA5AE9F24117C4B1FE649286651ECE45B3D"+
			"C2007CB8A163BF0598DA48361C55D39A69163FA8FD24CF5F"+
			"83655D23DCA3AD961C62F356208552BB9ED529077096966D"+
			"670C354E4ABC9804F1746C08CA237327FFFFFFFFFFFFFFFF", 16)
	pb := p.Bytes()
	m = make([]constbn.Base, (len(pb)/2)+2)
	constbn.Decode(m, pb)

	m0i = constbn.Ninv(m[1])

	mlen := (m[0] + 63) >> 5

	// Since we are currently testing the optimized implementation, we will give
	// it plenty of temporary space - 5 times the length of m which is the maximum window
	// size that can be used
	tmp = make([]constbn.Base, 5*mlen)
}

func randomInto(r io.Reader, b []byte) error {
	if _, err := io.ReadFull(r, b); err != nil {
		return errors.New("short random read")
	}
	return nil
}

//export prepare_inputs
func prepare_inputs(inputptr *C.uint8_t, classesptr *C.uint8_t) {
	allinputs := makeslice(inputptr, chunksize*measurements)
	classes := makeslice(classesptr, measurements)
	inputs := make([][]byte, measurements)
	for i := range inputs {
		inputs[i] = allinputs[i*chunksize : (i+1)*chunksize]
	}

	zerobn := []byte{0x0}

	var val []byte
	for i := range inputs {
		classes[i] = randombit()
		if classes[i] == 1 {
			val = make([]byte, 40)
			randomInto(rand.Reader, val)
		} else {
			val = zerobn
		}
		copy(inputs[i][0:40], val)
	}
	runtime.GC()
}

//export do_one_computation
func do_one_computation(dataptr *C.uint8_t) C.uint8_t {
	data := makeslice(dataptr, chunksize)

	result := make([]constbn.Base, 51)
	copy(result, x)
	result[0] = m[0]

	constbn.ModpowOpt(result, data[0:40], m, m0i, tmp)

	return (C.uint8_t)(uint8(1))
}

func main() {}
