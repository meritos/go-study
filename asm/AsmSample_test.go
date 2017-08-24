package asm

import (
	"testing"
	"fmt"
)


func TestAsmGames(t *testing.T) {


	z := AsmGames(100, 150)


	// vector1 := []int16{1, 2, 3, 4}
	// z := Sum(vector1)

	//fmt.Printf("\n\n\n  End Run ::: %d \n", z)
	fmt.Printf("\n %d, \n END\n", z)
}


func TestCPUCycles(t *testing.T){

	vector1 := make([]int16, 1000)
	
	a := StartTSCMeasure()

	// Benchmark partt
	Sum(vector1)

	b := StopTSCMeasure()

	cycleCounter := b - a
	if cycleCounter > 4000  {
		t.Errorf("\n ERROR ... took:  %d, expected: not pass 4 cycle for loop ", cycleCounter)
	} 


}