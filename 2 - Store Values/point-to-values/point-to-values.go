package main

import "fmt"

/*
	Whenever your program creates a container to store data, your computer allocates a space in memory at which to store the data
	An initialised variable consists of three parts
		- Name					"num"
		- Value					20
		- Memory Address		0x00009e068


	Go program supports the concept of "pointers" - a pointer can store the address of another variable & access the value stored at that address

	// Declaration of a pointer variable requires the type prefixxed by an * asterisx to indicate that this will be a pointer
	var ptr *int

	// The address of another variable can be assigned to the pointer by prefixxing the other variable's name with an & ampersand
	ptr = &num // Assigns an address to the pointer variable

	// The value at the assigned address can then be accessed by prefixxing the pointer name with an * asterisk
	*ptr // Points to the value at the assigned address
*/

func main() {
	var num int = 20
	var ptr *int = &num

	// Print the value & memory address of the regular integer variable
	fmt.Printf("num value: %v Type %T\n", num, num)
	fmt.Printf("num address: %v Type %T \n\n", ptr, ptr)

	// Print the referenced value & memory address of the pointer variable
	fmt.Printf("num via pointer: %v type: %T\n", *ptr, *ptr)
	fmt.Printf("ptr address: %v type: %T\n\n", &ptr, &ptr)

	// Statements to change the value stored in the integer variable - by assignment to the pointer variable
	*ptr = 100
	fmt.Printf("new num value: %v type: %T\n", num, num)

}
