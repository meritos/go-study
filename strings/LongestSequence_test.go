
package strings

import (

	"testing"
)


func TestLongestSequence_Errors(t *testing.T) {

	var wrongInputs = []string { 
		
		"[123", 
		"[]", 
		"23423,4  , 4]", 
		"[123, abc]", 
	    "[98078979, q]",
	    "[98078979,   ,]",
	}

	for _, wrongInput := range wrongInputs {

		_, _, err := LongestSequence(wrongInput)	

		if (err == nil) {

			t.Errorf("\n ERROR ... out:  %q, expected: %q", err, "err != null")
		}
	}

}


func TestLongestSequence_Suite(t *testing.T){

	  var tableTests = []struct {
	  		
	  		input          string 
	  		exptectedCount int 
	  		exptectedNum   string 
	  }{
	  	{"[123]", 3, "123"},
	  	{"[56789013]", 7, "56789013"},
	  	{"[56789013, 123, 1234]", 7, "56789013"},
	  	{"[123, 1234]", 4, "1234"},
	  	{"[123, 1234, 987654321]", 9, "987654321"},
	  	{"[0000234,   12314, 123, 56789013, 987]", 7, "56789013"},
	  }


	  for _, test := range tableTests {

		count, num, _ := LongestSequence(test.input)

		if count != test.exptectedCount {

			t.Errorf(" exptected: %d, result: %d \n", test.exptectedCount, count);
		} 

		if num != test.exptectedNum {

			t.Errorf(" exptected: %s, result: %s \n", test.exptectedNum, num);			
		}

	  }
}
