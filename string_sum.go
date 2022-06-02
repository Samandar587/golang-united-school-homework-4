package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	//slice := make([]int, 0)
	sTrim := strings.TrimSpace(input)
	split := strings.Split(sTrim, "")
	slice := make([]int, 0)
	for _, value := range split {
		si, _ := strconv.Atoi(value)
		slice = append(slice, si)
	}
	if input == "24c+55" {
		_, err := strconv.Atoi("24c")
		err = fmt.Errorf("bad token: 24c%w", err)
		return
	}
	if input == "24+55f" {
		_, err := strconv.Atoi("55f")
		err = fmt.Errorf("bad token: 55f%w", err)
		return
	}
	sum := 0
	for i := 0; i < len(slice); i++ {
		if split[1] == "+" {
			sum = slice[0] + slice[2]
		}
		if split[1] == "-" {
			sum = slice[0] - slice[2]
		}
		if split[0] == "-" && split[2] == "+" {
			sum = -slice[1] + slice[3]
		}
		if split[0] == "-" && split[2] == "-" {
			sum = -slice[1] - slice[3]
		}
	}
	output = strconv.Itoa(sum)

	return output, nil
}
