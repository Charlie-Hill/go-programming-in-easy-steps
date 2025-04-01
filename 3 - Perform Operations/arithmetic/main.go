package main

import "fmt"

/*

	The arithmetic operators commonly used in Go programs below:
	+ Addition
	- Subtraction
	* Multiplication
	/ Division
	% Remainder
	++ Increment
	-- Decrement

	Care must be taken to group expressions where more than one operator is used
	Operations within innermost () parentheses are performed first:
		a = b * c - d %e / f; - THIS IS UNCLEAR
		a = (b * c) - ( (d % e) / f ) - THIS IS CLEARER

*/

func main() {
	a := 8
	b := 4

	fmt.Println("Addition:\t", (a + b))
	fmt.Println("Subtraction:\t", (a - b))
	fmt.Println("Multiplication:\t", (a * b))
	fmt.Println("Division:\t", (a / b))
	fmt.Println("Remainder:\t", (a % b))

	a++
	fmt.Println("Increment:\t", a)
	b--
	fmt.Println("Decrement:\t", b)
}
