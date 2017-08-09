package strings

import (

	"fmt"
	"strings"
	"regexp"
)


/**
 * LongestSequence(in string) string
 *    
 *   given string of numbers calculage the consequence sequence  of digits
 *    [123, 1234, 567890] => result: 6
 *
 *  @param in string - list of numbers in form [23234, 234234, ....]
 *
 *  @return int       - the largest number of consequenced digits  
 *          resultNum - the number in the list that has the most consequence digits
 *          err       - used for input true form validation
 *    
 */
func LongestSequence(in string) (result int, resultNum string, err error) {

	// validate strings to be in form: [num1, num2, num3]
	valid, _ := regexp.MatchString("\\[((\\s)*([0-9])+(\\s)*)+((\\s)*(,)+(\\s)*([0-9])+(\\s)*)*\\]", in)
	if !valid {

		return 0, "", fmt.Errorf("Incorrect input format, expected: [num1, num2, ...]")
	}

	in = strings.Replace(in, "[", "", -1)
	in = strings.Replace(in, "]", "", -1)
	in = strings.Replace(in, " ", "", -1)
	nums  := strings.Split(in, ",")

	maxConDigits := 0
	maxConNum    := "0"
	for _, num := range nums {

		conDigits := countConsequence(num)

		if conDigits > maxConDigits {

			maxConDigits = conDigits
			maxConNum = num
		}
	}

	return maxConDigits, maxConNum, nil
}


/**
 * countConsequence(num string) int
 *
 *   The program check the digits in the input 
 *   num and returns consequence number of digits
 *     
 *     e.g: 123487 => return: 4   
 *
 *   @param num string - number to check
 *
 *   returns: int    - counted consequencence orderd digits in the number
 */
func countConsequence(num string) int {

	type direction int;
	const ( 

		INCREASE direction = 0
		DECREASE direction = 1
		RANDOM   direction = 2
	)

	var numDirection direction = RANDOM
	    maxConsequence   := 0
	    countConsequence := 0

	for i, _ := range num {

		if i+1 < len(num) && conseqInc(num[i], num[i+1]) && numDirection != INCREASE {

			numDirection = INCREASE

			if countConsequence > maxConsequence {
				maxConsequence = countConsequence
			}

			countConsequence = 1

			continue
		}

		if i+1 < len(num) && conseqInc(num[i], num[i+1]) && numDirection == INCREASE {

			numDirection = INCREASE
			
			countConsequence++

			if countConsequence > maxConsequence {
				maxConsequence = countConsequence
			}

			continue
		}

		if i+1 < len(num) && conseqDec(num[i], num[i+1]) && numDirection != DECREASE {

			numDirection = DECREASE
			countConsequence = 1

			continue
		}

		if i+1 < len(num) && conseqDec(num[i], num[i+1]) && numDirection == DECREASE {

			numDirection = DECREASE

			countConsequence++
			
			if countConsequence > maxConsequence {
				maxConsequence = countConsequence
			}

			continue
		}

		if i+1 < len(num) && numDirection != RANDOM {

			numDirection = RANDOM

			countConsequence = 0

			continue
		}
	}
	return maxConsequence + 1
}

/**
 *  conseqInc(num1 byte, num2 byte) bool
 *   
 *   @param num1 - single number
 *   @param num2 - single number
 *   returns true if num2 is folowing num1
 *              e.g: 5 and 6
 */
func conseqInc(num1 byte, num2 byte) bool {

	if num1 + 1 == num2 || (num1 == 57 && num2 == 48) {

		return true
	} else {

		return false
	}
}

/**
 *  conseqDec(num1 byte, num2 byte) bool
 *   
 *   @param num1 - single number
 *   @param num2 - single number
 *   returns true if num2 is folowing num1
 *              e.g: 5 and 4
 */
func conseqDec(num1 byte, num2 byte) bool {

	if num1 - 1 == num2 || (num1 == 48 && num2 == 57) {

		return true
	} else { 

		return false
	}
}

