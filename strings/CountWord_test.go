
package strings

import (

	"testing"
)


func TestCountWord_1(t *testing.T){

	expected := "Hello5 World5 Count5 My2 Words5 Long4 words5 also4 like4 Ronaldinhioooo14"
	out := CountWord("Hello World Count My Words Long words also like Ronaldinhioooo")

	if out != expected {

		t.Errorf("\n ERROR ... out:  %q, expected: %q", out, expected)
	}
}



func TestCountWord_2(t *testing.T){

	expected := ""
	out := CountWord("")

	if out != expected {

		t.Errorf("\n ERROR ... out:  %q, expected: %q", out, expected)
	} 
}

func TestCountWord_3(t *testing.T){

	expected := "oneWord7"
	out := CountWord("oneWord")

	if out != expected {

		t.Errorf("\n ERROR ... out:  %q, expected: %q", out, expected)
	} 
}

func TestCountWord_4(t *testing.T){

	expected := "Ronaldinhioooo14"
	out := CountWord("Ronaldinhioooo")

	if out != expected {

		t.Errorf("\n ERROR ... out:  %q, expected: %q", out, expected)
	} 
}


func BenchmarkCountWord_4(b *testing.B){

	b.ReportAllocs()

	for i:=0; i < 1000000; i++ {

		CountWord("Hello World Count My Words Long words also like Ronaldinhioooo ")
	}



}




