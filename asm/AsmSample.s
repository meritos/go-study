
// TODO: study memory dumps
// TODO: study DATA for arrays
// TODO: study flags


#include "textflag.h"
#include "funcdata.h"

//
// Define: TEXT ·Add(SB), [$stack_size-arg_size] 
// 
// Stack size defined: 
//  (!) ... todo: acording vars you use or args you pass to calls 
//
// Arguments defined as 24 bytes
//   x, y => 0-16  : 2 int64
//   ret  => 16-24  : 1 int64
TEXT ·AsmGames(SB), NOSPLIT|NOFRAME, $320-24
	
	// If you have intention to 
	// do calls from this ASM code
	// you better use this makro, 
	// for more info: https://golang.org/doc/asm#runtime
	NO_LOCAL_POINTERS	

	
	// Init args for the CALL
	// make CALL with the args 
	MOVQ	SP,    0(SP)
	MOVQ    $1,    8(SP)
	MOVQ    $1,   16(SP)
	MOVQ    $1,   24(SP)
	CALL    ·DebugInfo(SB)
	// check what (TLS) contains
    //CALL ·PrintStack+0(SB)

    // Example of printing value
    MOVQ $125, R8
    SHRQ $1, R8
	//MOVQ	R8,    0(SP)
	//CALL    ·PrintVal(SB)

	LFENCE
	RDTSC
	SHLQ  $32,  DX

	MOVQ  AX,   BX
	ADDQ  DX,   BX

	// Benchmark code
  //  PUSHFQ
  //  POPQ R9	
//******
    MOVQ $1000,   R10

loop:
	MOVQ  AX,   R11
	MOVQ  AX,   R12
	MOVQ  AX,   R13
	MOVQ  AX,   R14

	SUBQ $1,    R10
JNZ loop	
//******


	RDTSC
	SHLQ  $32,  DX
	MOVQ  AX,   48(SP)
	ADDQ  DX,   48(SP)

	MOVQ  BX,   32(SP)

	MOVQ	48(SP), R9
	MOVQ	R9, 0(SP)
	CALL ·PrintVal(SB)

	MOVQ	32(SP), R9
	MOVQ	R9, 0(SP)
	CALL ·PrintVal(SB)
		

	MOVQ	32(SP), R9
	MOVQ	R9, 0(SP)

	MOVQ	48(SP), R9
	MOVQ	R9, 8(SP)
	CALL ·PrintDif(SB)

	LFENCE




	// Example of 
    PUSHFQ
    POPQ R9
	MOVQ	R9,    0(SP)
	CALL    ·PrintFlags(SB)


	// * Actual x+y function *
	MOVQ x+0(FP), BX
	MOVQ y+8(FP), BP
	ADDQ BP, BX
	MOVQ BX, ret+16(FP)
	RET


//
// func StartTSCMeasure unit64
//
//   Retrieve the CPU cycles counter (TSC)
// 
//   $0-8 : No Stack been used - 8 bytes for 
//          return int64
//
TEXT ·StartTSCMeasure(SB), NOSPLIT, $0-8

	// Flush all in progress instructions
	LFENCE 

	// Retrieve cycles counter
	RDTSC

	// Combine DX:AX
	SHLQ  $32,  DX
	MOVQ  AX,   BX
	ADDQ  DX,   BX

	MOVQ BX, ret+0(FP)
	RET

//
// func StopTSCMeasure unit64
//
//   Retrieve the CPU cycles counter (TSC)
// 
//   $0-8 : No Stack been used - 8 bytes for 
//          uint64 return
//
TEXT ·StopTSCMeasure(SB), NOSPLIT, $0-8

	// Retrieve cycles counter
	RDTSC

	// Combine DX:AX
	SHLQ  $32,  DX
	MOVQ  AX,   BX
	ADDQ  DX,   BX

	MOVQ BX, ret+0(FP)

	// Releae LFENCE
	LFENCE 
	RET


//
// func Add(x, y int64) int64 
//
// Stack size defined: $0 
//
// Arguments defined as 24 bytes
//   x, y => 0-16  : 2 int64
//   ret  => 16-24  : 1 int64
//
TEXT ·Add(SB), NOSPLIT, $0-24

	MOVQ x+0(FP), BX
	MOVQ y+8(FP), BP
	ADDQ BP, BX
	MOVQ BX, ret+16(FP)
	RET



//
// func Sum(v []int16) int16
//
// Stack size defined: $0 
//
// ... todo: comment array param
// 
TEXT ·Sum(SB), NOSPLIT, $0-26
	MOVQ v_base+0(FP), R8
	MOVQ v_len+8(FP), R9
	SHLQ $1, R9
	ADDQ R8, R9
	MOVQ $0, R10

loop:
	CMPQ R8, R9
	JE   end
	ADDW (R8), R10
	ADDQ $2, R8
	JMP  loop

end:
	MOVW R10, ret+24(FP)
	RET
