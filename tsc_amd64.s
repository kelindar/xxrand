#include "textflag.h"

// func x64tsc() uint64
TEXT ·x64tsc(SB),NOSPLIT,$0-8
	MFENCE
	LFENCE
	RDTSCP
	SHLQ	$32, DX
	ADDQ	DX, AX
	MOVQ	AX, ret+0(FP)
	RET
