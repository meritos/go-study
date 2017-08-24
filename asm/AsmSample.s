

// TODO: study memory dumps
// TODO: study DATA for arrays 
// TODO: study func address and memory 


#include "textflag.h"
#include "funcdata.h"

//
// Local static array of ints 
//
DATA intArray<>+0x00(SB)/8, $0
DATA intArray<>+0x08(SB)/8, $1
DATA intArray<>+0x10(SB)/8, $2
DATA intArray<>+0x18(SB)/8, $3
DATA intArray<>+0x20(SB)/8, $4
DATA intArray<>+0x28(SB)/8, $5
DATA intArray<>+0x30(SB)/8, $6
DATA intArray<>+0x38(SB)/8, $7
DATA intArray<>+0x40(SB)/8, $8
GLOBL intArray<>(SB), (RODATA | NOPTR), $0x48


DATA intArray2+0x00(SB)/8, $0xFF
DATA intArray2+0x08(SB)/8, $0xEE
DATA intArray2+0x10(SB)/8, $0xDD
DATA intArray2+0x18(SB)/8, $3
DATA intArray2+0x20(SB)/8, $4
DATA intArray2+0x28(SB)/8, $5
DATA intArray2+0x30(SB)/8, $6
DATA intArray2+0x38(SB)/8, $7
DATA intArray2+0x40(SB)/8, $8
GLOBL intArray2(SB), (RODATA | NOPTR), $0x48

//
// Define: TEXT ·Add(SB), [$stack_size-arg_size] 
// 
// Stack size defined: 
//    $320 - to be able to play with bunch of params
//           without breaking anything, practice 
//           for samples only  
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


	BYTE $0x90

	BYTE $0x91
	BYTE $0x91
	BYTE $0x91
	BYTE $0x91

	BYTE $0x90
	BYTE $0x90
	BYTE $0x90
	BYTE $0x90

	MOVQ	$20, R9
	MOVQ	R9, 0(SP)
	CALL ·PrintVal(SB)


	// Example of 
    PUSHFQ
    POPQ    R9
	MOVQ	R9,    0(SP)
	CALL    ·PrintFlags(SB)


	LEAQ intArray<>(SB), DX
	// Init args for the CALL
	// make CALL with the args 
	MOVQ	DX,    0(SP)
	MOVQ    $10,    8(SP)
	CALL    ·PrintMem(SB)

	LEAQ intArray2(SB), DX
	// Init args for the CALL
	// make CALL with the args 
	MOVQ	DX,    0(SP)
	MOVQ    $10,    8(SP)
	CALL    ·PrintMem(SB)

	LEAQ ·AsmGames(SB), DX
	// Init args for the CALL
	// make CALL with the args 
	MOVQ	DX,    0(SP)
	MOVQ    $10,    8(SP)
	CALL    ·PrintMem(SB)
	

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

// 
// func BubbleSort(v []int64)  
//
//  The function will recive unsorted 
//  array of uint64 and sort it using 
//  bubble sort algorithm
// 
// Stack size defined: $0 
//
// Arguments defined as 16 bytes
//   vector address => 0-8  : 1 int64
//   vector length  => 8-16  : 1 int64
//
TEXT ·BubbleSortAsm(SB), NOSPLIT, $0-16

	// Load arguments
	MOVQ v_base+0(FP), R8
	MOVQ v_len+8(FP), R9

	// input validation
	CMPQ R9, $1
	JLE  not_valid_input 	

	swapped_run_again:
		
		MOVQ $0, R15 // R15 will indicatae the swap 
		MOVQ R8, R11 // R11 will be the loop counter
		
		MOVQ R9, R12 // R12 will be the end loop marker
		SHLQ $3, R12 // lenght * 8 (sizeof(int64))
		ADDQ R8, R12 // add the array start
		SUBQ $8, R12 // finish on the len-1 element

	loop: 

		// check for swap
		// compare (R11), (R11)(8)
		// if (R11) is greater swap 
		MOVQ  (R11), R13
		MOVQ 8(R11), R14
		CMPQ R13, R14
		JLE  no_swap 

		// Swap part
		MOVQ $1,  R15 // Flag the swap happened
		MOVQ R13, 8(R11)
		MOVQ R14,  (R11)
	no_swap:

		ADDQ $8, R11 // Increment by sizeof(int64)
		CMPQ R11, R12
		JNE loop

		CMPQ R15, $0
		JNZ swapped_run_again

	not_valid_input:
		
	RET

