package asm

import (
	"testing"
	"fmt"
)


func TestAsmSum(t *testing.T) {


	z := Add(100, 150)
	fmt.Printf("\n\n\n  Wow !!! %d \n", z)
}