// +build !gccgo

#include "textflag.h"
#include "funcdata.h"

TEXT ·Add(SB),NOSPLIT,$0
	
	// Invalid memory address
	// CALL fmt·Printf(SB)

	MOVQ x+0(FP), BX
	MOVQ y+8(FP), BP

	// LEAL	runtime·emptystring(SB), SI
	// CALL runtime·exit(SB)

	// CALL	syscall·errstr(SB)

	ADDQ BP, BX
	MOVQ BX, ret+16(FP)

	// CALL runtime·entersyscall(SB)
	// CALL runtime·exitsyscall(SB)

	CALL ·PrintStack+0(SB)

	RET
	