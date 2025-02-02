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
		si, err := strconv.Atoi(value)
		if value != "+" && value != "-" {
			if err != nil {
				return "", fmt.Errorf("bad token inserteed: %w", err)
			}
		}

		slice = append(slice, si)
	}
	fmt.Println(split)
	fmt.Println(slice)
	lenSlice := len(slice)
	if sTrim == "" {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}
	if lenSlice < 3 || lenSlice > 4 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
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
