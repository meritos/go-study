
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


import (
	
	"fmt"
	"runtime/debug"
	"unsafe"
	"encoding/binary"
)


func Add(x, y int64) int64


func PrintStack(sb int64, sb2 int64){

	fmt.Printf("\n\n = ALL GOOD =\n")
	debug.PrintStack()
	fmt.Printf(" = ALL GOOD = \n\n")

	// Stack Base
	fmt.Printf("\nSB: 0x%x \n\n", sb)
	
	// ptr := unsafe.Pointer(&sb)
	// myPtr uintptr = &sb;

	for i := 0; i < 20; i=i+2{

		// Prefix line
		fmt.Printf(" 0x%016x: ", sb + int64(i) * 8)

		// Memory Data
		ptr := (*uint64)(unsafe.Pointer(uintptr(sb + int64(i * 8) )))	
		fmt.Printf("%016x ", *(ptr) )

		ptr2 := (*uint64)(unsafe.Pointer(uintptr(sb + int64((i + 1) * 8) )))	
		fmt.Printf("%016x ", *(ptr2) )

		// Suffix line
		bs := make([]byte, 8)
		binary.LittleEndian.PutUint64(bs, *(ptr))
		
		bs2 := make([]byte, 8)
		binary.LittleEndian.PutUint64(bs2, *(ptr2))
		
		fmt.Printf(" .... ")
		fmt.Printf("%q%q%q%q%q%q%q%q ", bs[7], bs[6], bs[5], bs[4], bs[3], bs[2], bs[1], bs[0]  )
		fmt.Printf("%q%q%q%q%q%q%q%q", bs2[7], bs2[6], bs2[5], bs2[4], bs2[3], bs2[2], bs2[1], bs2[0]  )
		fmt.Println()
		

	}	

    
}






