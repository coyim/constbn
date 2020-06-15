all: dut_aes32 dut_aesbitsliced dut_cmpmemcmp dut_cmpct dut_donna dut_donnabad
go: dut_go
constbn: dut_constbn dut_constbn_opt

OBJS = src/cpucycles.o src/fixture.o src/random.o \
src/ttest.o src/percentile.o
OBJS_AES32 = dut/aes32/rijndael-alg-fst.o
OBJS_DONNA = dut/donna/curve25519-donna.o
OBJS_DONNABAD = dut/donnabad/curve25519-donnabad.o
OBJS_AESBITSLICED = dut/aesbitsliced/afternm_aes128ctr.o \
dut/aesbitsliced/beforenm_aes128ctr.o \
dut/aesbitsliced/common_aes128ctr.o \
dut/aesbitsliced/consts_aes128ctr.o \
dut/aesbitsliced/int128_aes128ctr.o \
dut/aesbitsliced/stream_aes128ctr.o \
dut/aesbitsliced/xor_afternm_aes128ctr.o
SHRDOBJ_GO = dut/go/dut_go.so
SHRDOBJ_GO_SRC := $(shell find dut/go -type f -name '*.go')
SHRDOBJ_GO_CONSTBN = dut/constbn/dut_constbn.so
SHRDOBJ_GO_CONSTBN_SRC := $(shell find dut/constbn -type f -name '*.go')
SHRDOBJ_GO_CONSTBN_OPT = dut/constbn_opt/dut_constbn.so
SHRDOBJ_GO_CONSTBN_OPT_SRC := $(shell find dut/constbn_opt -type f -name '*.go')

CC=clang
OPTIMIZATION=-O2
#CFLAGS	= -Weverything -O0 -fsanitize=memory -fno-omit-frame-pointer -g 
CFLAGS	= $(OPTIMIZATION)
LIBS	= -lm
#LDFLAGS	= -fsanitize=memory -fno-omit-frame-pointer -g 
#LDFLAGS = -Weverything $(OPTIMIZATION)
LDFLAGS = $(OPTIMIZATION)

INCS	= -Iinc/

dut_aes32: $(OBJS) $(OBJS_AES32) dut/aes32/dut_aes32.c
	$(CC) $(LDFLAGS) $(INCS) -o dudect_aes32_$(OPTIMIZATION) dut/aes32/$@.c $(OBJS) $(OBJS_AES32) $(LIBS)

dut_aesbitsliced: $(OBJS) $(OBJS_AESBITSLICED) dut/aesbitsliced/dut_aesbitsliced.c
	$(CC) $(LDFLAGS) $(INCS) -o dudect_aesbitsliced_$(OPTIMIZATION) dut/aesbitsliced/$@.c $(OBJS) $(OBJS_AESBITSLICED) $(LIBS)

dut_cmpmemcmp: $(OBJS) dut/cmpmemcmp/dut_cmpmemcmp.c
	$(CC) $(LDFLAGS) $(INCS) -o dudect_cmpmemcmp_$(OPTIMIZATION) dut/cmpmemcmp/$@.c $(OBJS) $(LIBS)

dut_cmpct: $(OBJS) dut/cmpct/dut_cmpct.c
	$(CC) $(LDFLAGS) $(INCS) -o dudect_cmpct_$(OPTIMIZATION) dut/cmpct/$@.c $(OBJS) $(LIBS)

dut_donna: $(OBJS) $(OBJS_DONNA) dut/donna/dut_donna.c
	$(CC) $(LDFLAGS) $(INCS) -o dudect_donna_$(OPTIMIZATION) dut/donna/$@.c $(OBJS) $(OBJS_DONNA) $(LIBS)

dut_donnabad: $(OBJS) $(OBJS_DONNABAD) dut/donnabad/dut_donnabad.c
	$(CC) $(LDFLAGS) $(INCS) -o dudect_donnabad_$(OPTIMIZATION) dut/donnabad/$@.c $(OBJS) $(OBJS_DONNABAD) $(LIBS)

dut_go: $(OBJS) $(SHRDOBJ_GO) dut/go/dut_go.c
	$(CC) $(LDFLAGS) $(INCS) -o dudect_go_$(OPTIMIZATION) dut/go/$@.c $(OBJS) $(SHRDOBJ_GO) $(LIBS)

$(SHRDOBJ_GO): $(SHRDOBJ_GO_SRC)
	go build -o $@ -buildmode=c-shared $(SHRDOBJ_GO_SRC)

dut_constbn: $(OBJS) $(SHRDOBJ_GO_CONSTBN) dut/constbn/dut_constbn.c
	$(CC) $(LDFLAGS) $(INCS) -o dudect_constbn_$(OPTIMIZATION) dut/constbn/$@.c $(OBJS) $(SHRDOBJ_GO_CONSTBN) $(LIBS)

dut_constbn_opt: $(OBJS) $(SHRDOBJ_GO_CONSTBN_OPT) dut/constbn_opt/dut_constbn.c
	$(CC) $(LDFLAGS) $(INCS) -o dudect_constbn_opt_$(OPTIMIZATION) dut/constbn_opt/dut_constbn.c $(OBJS) $(SHRDOBJ_GO_CONSTBN_OPT) $(LIBS)

$(SHRDOBJ_GO_CONSTBN): $(SHRDOBJ_GO_CONSTBN_SRC)
	go build -o $@ -buildmode=c-shared $(SHRDOBJ_GO_CONSTBN_SRC)

$(SHRDOBJ_GO_CONSTBN_OPT): $(SHRDOBJ_GO_CONSTBN_OPT_SRC)
	go build -o $@ -buildmode=c-shared $(SHRDOBJ_GO_CONSTBN_OPT_SRC)

.c.o:
	$(CC) $(CFLAGS) $(INCS) -c $< -o $@

clean:
	rm -f $(OBJS) $(OBJS_AES32) $(OBJS_AESBITSLICED) $(OBJS_DONNA) $(OBJS_DONNABAD) $(SHRDOBJ_GO) dudect_* *.exe a.out
