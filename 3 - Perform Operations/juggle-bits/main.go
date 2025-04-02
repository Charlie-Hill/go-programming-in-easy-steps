package main

import "fmt"

/*

	Each byte comprises eight bits that can each contain a 1 or 0 to store a binary number representing a decimal value from 0 to 255
	Each bit contributes a decimal component only when that bit contains a 1
	Components are designated right-to-left from the "Least Significant Bit" (LSB) to the "Most Significant Bit" (MSB)The binary number in the bit pattern 00110010 represents the decimal number 50 (2 + 16 + 32)

	It is possible to manipulate individual parts of a byte using the Go "bitwise operators" listed below
	& AND 0011 & 0101 = 0001
	| OR  0011 | 0101 = 0111
	^ XOR 0011 ^ 0101 = 0111
	^ NOT ^ 0101 = 1010
	&^ AND NOT 0011 &^ 0101 = 0010
	<< Shift left 0010 << 2 = 1000
	>> Shift right 1000 >> 2 = 0010

*/

func main() {
	var flags byte = 10 // Bonary 1010 (1*x 0x4 1x2 0x1)

	fmt.Printf("Flag #1: %v\n", (flags&1) > 0)
	fmt.Printf("Flag #2: %v\n", (flags&2) > 0)
	fmt.Printf("Flag #3: %v\n", (flags&3) > 0)
	fmt.Printf("Flag #4: %v\n", (flags&4) > 0)

	flags = flags >> 1
	fmt.Printf("New Value: %08b\t%v\n", flags, flags)
}
