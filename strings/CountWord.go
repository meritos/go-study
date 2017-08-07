package strings

import (

	"fmt"
	"strconv"
	"strings"
)


func cheet(){
	fmt.Print("");
}

/**
 * CountWord2(in string) string
 *
 *		...
 */
func CountWord(in string) string {

	in = strings.Trim(in, " ")

	words  := strings.Split(in, " ")
	result := make([]byte, len(in) + calcAddition(in))


	copyIdx := 0;
	for _, word := range words {

		// *** ver 1 ***
		//newWord := strings.Join([]string{word, strconv.Itoa(len(word)), " "}, "")
		//copy(result[copyIdx:], newWord)
		//copyIdx += len(newWord)

		// *** ver 2 ***
		copy(result[copyIdx:], word)
		copyIdx += len(word)

		wordLenStr := strconv.Itoa(len(word))
		copy(result[copyIdx:], wordLenStr)
		copyIdx += len(wordLenStr)

		if copyIdx + 1 < len(result){

			 copy(result[copyIdx:], " ")
			 copyIdx++
		}

	}

	return string(result)

}


/**
 * calcAddition(in string) int
 *
 *     ...
 *
 */
func calcAddition(in string) int {

		// counts spaces in the string 
		bufferSizeAdd := 0;
		wordLenCount  := 0;
		for idx, elem := range in {

			if elem == 32 || idx == len(in) - 1 {

				bufferSizeAdd++

				if wordLenCount > 9 {
					
					bufferSizeAdd++
				} else if wordLenCount > 99 {

					bufferSizeAdd = bufferSizeAdd + 2;
				}

				wordLenCount = 0
			}

			wordLenCount++
		}

		return bufferSizeAdd;
}

