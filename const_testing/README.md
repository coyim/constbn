# Constant time testing of the constbn library

In this directory we have collected the information and tools necessary to make it possible to test the constant time
nature of the implementations in `constbn`. For now, the only implementations that have been tested are the `ModpowInt`
and `ModpowOpt` implementations - and only checking against constant time on the exponent, since that's the most
important factor in many cryptographic implementations. However, since these two implementations use most of the rest of
the functionality in the library, we can say with high certainty that those implementations are also constant time.

This directory contains two directories that contain the implementations for the tests, a Makefile, and the results of
output. In order to run these tests, you need to clone the tool Dudect. There is a version of this tool that has added
support for Golang integration. You can find it at https://github.com/ansemjo/dudect - check it out at and copy the
Makefile from this directory into the base directory. Also copy the two directories from here into the `dut`
directory. Assuming you have the `constbn` project in your `GOPATH` you should be able to run `make constbn` and this
will generate two executables: `dudect_constbn_-02` and `dudect_constbn_opt_-02`. When you run one of these, it will
start trying to disprope the constant time nature of the `constbn` code.


## Results

The two files `result_constbn_modpow` and `result_constbn_modpow_opt` contain the output of running the above commands
for roughly 7 days. The important part is the `max t` value, which should never go over 10. As you can see, for the 96
hours, these values almost always hover between 1 and 2, which is a very good indiciation that these implementations are
in fact constant time.


## Attribution

This code is lightly modified from code originally created by Stefan Marsiske, ROS. 
