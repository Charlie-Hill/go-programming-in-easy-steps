/*
PrintF Format Specifiers:
	%s - A string of Characters (i.e. "Go Fun!")
	%d - An integer (-32768 to +32767) (i.e. 100)
	%f - A floating-point number (i.e. 0.123456)
	%c - A single character (i.e 'A')
	%t - A boolean value (i.e. true)
	%p - A machine memory address (i.e. 0x0022FF34)
	%v - The value in a default format (any of the above)
	%T - The datatype of the varialbe (i.e. int)
*/

package main

import "fmt"

func main() {
	num := 100
	pi := 3.1415926536

	fmt.Printf("num %v type %T\n", num, num)
	fmt.Printf("num %v type %T\n\n", pi, pi)

	fmt.Printf("%%7d displays %7d\n", num)
	fmt.Printf("%%07d displays %07d\n\n", num)

	fmt.Printf("Pi is approximately %1.10f\n", pi)
	fmt.Printf("Right-aligned %20.3f rounded pi\n", pi)
	fmt.Printf("Left aligned %-20.3f rounded pi\n\n", pi)
}
