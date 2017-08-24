
package asm

//
//  Study links: 
//
//    https://9p.io/sys/doc/asm.html
//
//    https://goroutines.com/asm
//    https://blog.sgmansfield.com/2017/04/a-foray-into-go-assembly-programming/
//    https://fosdem.org/2017/schedule/event/go_scaling/
//
//    https://medium.com/golangspec/making-debugger-for-golang-part-i-53124284b7c8
//    http://www.agner.org/optimize/
//
//    https://blog.altoros.com/golang-part-1-main-concepts-and-project-structure.html
//    https://go.godbolt.org/

//   https://en.wikibooks.org/wiki/X86_Assembly/X86_Architecture


import (
	
	"fmt"
	"unsafe"
	"encoding/binary"
)

func AsmGames(x, y int64) int64

func Add(x, y int64) int64
func Sum(v []int16) int16
func BubbleSortAsm(v []int64)  

func StartTSCMeasure() int64
func StopTSCMeasure() int64


func PrintVal(val int64){
	fmt.Printf("\n val: 0x%016x | %10d | %064b |\n", val, val, val)
}

func PrintDif(val1, val2 int64){
	fmt.Printf("\n diff : %d - %d = %d |\n", val2, val1, (val2 - val1))
}


//
// Print flags assuming the int64 
// is the flags register passed
// source: https://en.wikipedia.org/wiki/FLAGS_register
//   
//  @param  flags int64 - flags register
//   usually obtained by PUSHFQ, POPFQ 
//
func PrintFlags(flags int64){

	fmt.Println();
	fmt.Printf("[ 0] CF - %d\n", flag(flags & 1))
	fmt.Printf("[ 1]    - %d\n", flag(flags & 2))
	fmt.Printf("[ 2] PF - %d\n", flag(flags & (2<<1)    ))
	fmt.Printf("[ 3]    - %d\n", flag(flags & (2<<2)    ))
	fmt.Printf("[ 4] AF - %d\n", flag(flags & (2<<3)    ))
	
	fmt.Printf("[ 5]    - %d\n", flag(flags & (2<<4)    ))
	fmt.Printf("[ 6] ZF - %d\n", flag(flags & (2<<5)    ))
	fmt.Printf("[ 7] SF - %d\n", flag(flags & (2<<6)    ))
	fmt.Printf("[ 8] TF - %d\n", flag(flags & (2<<7)    ))
	fmt.Printf("[ 9] IF - %d\n", flag(flags & (2<<8)    ))
	fmt.Printf("[10] DF - %d\n", flag(flags & (2<<9)    ))
	fmt.Printf("[11] OF - %d\n", flag(flags & (2<<10)   ))
	fmt.Printf("[12] I1 - %d\n", flag(flags & (2<<11)   ))
	fmt.Printf("[13] I2 - %d\n", flag(flags & (2<<12)   ))
	fmt.Printf("[14] NT - %d\n", flag(flags & (2<<13)   ))
	fmt.Printf("[15]    - %d\n", flag(flags & (2<<14)   ))

	fmt.Printf("[16] RF - %d\n", flag(flags & (2<<15)   ))
	fmt.Printf("[17] VM - %d\n", flag(flags & (2<<16)   ))
	fmt.Printf("[18] AC - %d\n", flag(flags & (2<<17)   ))
	fmt.Printf("[19] VF - %d\n", flag(flags & (2<<18)   ))
	fmt.Printf("[20] VP - %d\n", flag(flags & (2<<19)   ))
	fmt.Printf("[21] ID - %d\n", flag(flags & (2<<20)   ))
}


func flag(f int64) byte{
	if (f > 0){
		return 1
	} else {
		return 0
	}
}

func PrintMem(addr int64, size int){

	fmt.Printf("\n\n    === Memory From: [ 0x%x ] Size: [ %d x 8 bytes ] === \n", addr, size)

	for i := 0; i < 20; i=i+2{

		// Prefix line
		fmt.Printf(" 0x%016x: ", addr + int64(i) * 8)

		// Memory Data
		ptr := (*uint64)(unsafe.Pointer(uintptr(addr + int64(i * 8) )))	
		fmt.Printf("%016x ", *(ptr) )

		ptr2 := (*uint64)(unsafe.Pointer(uintptr(addr + int64((i + 1) * 8) )))	
		fmt.Printf("%016x ", *(ptr2) )

		// Suffix line
		bs := make([]byte, 8)
		binary.LittleEndian.PutUint64(bs, *(ptr))
		
		bs2 := make([]byte, 8)
		binary.LittleEndian.PutUint64(bs2, *(ptr2))
		


		fmt.Printf(" |  ")
		fmt.Printf("%c%c%c%c%c%c%c%c ", nice(bs[7]), nice(bs[6]), nice(bs[5]), nice(bs[4]), 
										nice(bs[3]), nice(bs[2]), nice(bs[1]), nice(bs[0])  )
		fmt.Printf("%c%c%c%c%c%c%c%c \n", nice(bs2[7]), nice(bs2[6]), nice(bs2[5]), nice(bs2[4]), 
										nice(bs2[3]), nice(bs2[2]), nice(bs2[1]), nice(bs2[0])  )
		

	}	
}

// Print only if it is 32-126 otherwise .
func nice(c byte) byte{
	if (c >= 32 && c < 126){
		return c
	} else {
		return 46
	}
}




func swap(arrayzor []int64, i, j int64) {
	tmp := arrayzor[j]
	arrayzor[j] = arrayzor[i]
	arrayzor[i] = tmp
}

func BubbleSort(arrayzor []int64) {

	swapped := true;
	for swapped {
		swapped = false
		var i int64
		for i = 0; i < int64(len(arrayzor)) - 1; i++ {
			if arrayzor[i + 1] < arrayzor[i] {
				swap(arrayzor, i, i + 1)
				swapped = true
			}
		}
	}	
}


