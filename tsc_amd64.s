#include "textflag.h"

// func next() uint64
TEXT ·next(SB),NOSPLIT,$0-8
	MFENCE
	LFENCE
	RDTSCP
	SHLQ	$32, DX
	ADDQ	DX, AX
	MOVQ	AX, ret+0(FP)
	RET
