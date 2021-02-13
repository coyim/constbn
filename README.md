# constbn - a constant time Golang BigNum library

[![Build Status](https://github.com/coyim/constbn/workflows/constbn%20CI/badge.svg)](https://github.com/coyim/constbn/actions?query=workflow%3A%22constbn+CI%22)
[![Coverage Status](https://coveralls.io/repos/coyim/constbn/badge.svg?branch=main&service=github)](https://coveralls.io/github/coyim/constbn?branch=main)

This is an implementation of bignums with a focus on constant time operations. Unless anything else is mentioned, all
operations are constant time. The initial implementation is based on the i31 implementation from BearSSL. It uses uint32
values as the limbs, but only 31 bits are actually used.

The main goal of this implementation is to make it possible to have a generic constant time modular exponentiation
operation on bignums large enough to implement modern cryptographyic algorithms, since the Golang big/int library is not
constant time. Other operations might be added with time, making this a more generic library, but the focus is initially
to serve the needs of the otr3 project.

## API Documentation

[![GoDoc](https://godoc.org/github.com/coyim/constbn?status.svg)](https://godoc.org/github.com/coyim/constbn)

## Security and assumptions

- The code in this library assumes that the uint32 multiplication routines are constant time on the machine in question.
- The constant time nature of the modpow (Exp) implementation has been tested using `dudect`. You can see the code and
  the results from this testing in the `const_testing` sub directory in this repository. The summary is that this
  specific tests gives a good inidication that the code is actually constant time.

## Caveats and notes

- This is not a general purpose bignum implementation. It is specifically aimed at cryptography, and specifically
  cryptography in a constant time setting, with specific limitiations. For example, the Int type does not support
  negative numbers. It is also not possible to do an exponentiation with an even modulus.
- The code has not been audited. That said, it is a fairly straightforward translation from a small subset of BearSSL,
  and I have added a significant amount of test vectors. If something is wrong with this implementation, it's a strong
  possibility that something is also wrong with the BearSSL implementation.
- The API is subject to change
- Currently, there is no good documentation - the code itself is the most appropriate to read to understand what's going on right now.
- Unless specifically documented, all operations are constant time.


## Authors

- Centro de Autonomía Digital


## License

This project is licensed under the GNU GENERAL PUBLIC LICENSE VERSION 3.0.
