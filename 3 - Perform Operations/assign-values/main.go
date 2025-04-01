package main

import "fmt"

/*

	The operators used in Go to assign values are listed below:
	Operator '='  | Example (a = b)  | Equivalent (a = b)
	Operator '+=' | Example (a += b) | Equivalent (a + b)
	Operator '-=' | Example (a -= b) | Equivalent (a - b)
	Operator '*=' | Example (a *= b) | Equivalent a = (a * b)
	Operator '/=' | Example (a /= b) | Equivalent a = (a / b)
	Operator '%=' | Example (a %= b) | Equivalent a = (a % b)

*/

func main() {

	var a, b int
	a, b = 8, 4

	fmt.Println("Assigned Values:\ta = ", a, "\tb = ", b)

	a += b
	fmt.Println("Add & Assign\t a = ", a)

	a -= b
	fmt.Println("Subtract & Assign\t a = ", a)

	a *= b
	fmt.Println("Multiply & Assign\t a = ", a)

	a /= b
	fmt.Println("Divide & Assign\t a = ", a)

	a %= b
	fmt.Println("Remainder Assigned:\t a = ", a)
}
