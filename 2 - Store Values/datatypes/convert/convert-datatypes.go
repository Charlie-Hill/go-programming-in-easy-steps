package main

import (
	"fmt"
	"strconv"
)

func main() {
	// The strconv.Atoi() function converts a string to an int
	num, err := strconv.Atoi("42")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%d (%T)\n", num, num)

	// The strconv.Itoa() function converts an int to a string (no err return)
	str := strconv.Itoa(num)

	fmt.Printf("%s (%T)\n", str, str)

	// To convert a string value to a float64 value, you can use strconv.ParseFloat(string, FloatSize)
	float, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%f (%T)\n", float, float)

	// Casting a conversion is simply a matter of stating the required type followed by parentheses containing the value to be converted
	var decimal float64 = float64(num)

	fmt.Printf("(Casted) %f (%T)\n", decimal, decimal)

	// When a variable of the byte data type (uint8 - an 8bit unsigned integer)
	// 	a single character will store a numerical value that is the ASCII character code for that character
	var char byte = 'A'
	fmt.Printf("%c (%T)\n", char, char)

	// Can easily be casted to a string
	var castedStr = string(char)
	fmt.Printf("(Casted) %s (%T)\n", castedStr, castedStr)

}
