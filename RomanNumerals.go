package main

import (
	"fmt"
)

func main() {

	var RomanNumeral string

	RomanNumeral = decimalToRoman(499)
	fmt.Println(RomanNumeral)

}

func decimalToRoman(decimal int) (roman string) {

	var retval string

	// set up string representations of parts of numbers
	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	hundreds := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	thousands := []string{"", "M", "MM", "MMM"}

	// get roman numeral by concatenating the parts indexed from the slices above
	retval = thousands[getDigitFromNumber(1000, decimal)] + hundreds[getDigitFromNumber(100, decimal)] + tens[getDigitFromNumber(10, decimal)] + ones[getDigitFromNumber(1, decimal)]
	return retval
}

// getDigitFromNumber extracts the value of digit from decimal, so if digit were 100 and decimal 2145 this will return 1
func getDigitFromNumber(digit int, decimal int) int {

	var retval int

	// first get rid of higher units, so for instance, if we are looking at 10s we lose 1000s and 100s here
	// so 2145 becomes 45
	retval = decimal % (digit * 10)

	// now we have to truncate the earlier digits so using the example above 45 now becomes 4
	retval = retval / digit

	return retval
}
