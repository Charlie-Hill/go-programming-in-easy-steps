package main

import "fmt"

func main() {
	example()

	const pi = 3.14159
	const (
		red = iota + 1
		yellow
		green
		brown
		blue
		pink
		black
	)

	fmt.Printf("Pi approximately: %v\n\n", pi)

	fmt.Printf("Red: %v point \n", red)
	fmt.Printf("Blue: %v point \n", blue)
	fmt.Printf("Black: %v point \n", black)
}

func example() {
	const oneMillion = 1000000

	// Multiple constants of the same or different types can be created as a single declaration
	const twoMillion, twoThousaand = 2000000, "2000"

	// A constant declaration can employ a "constant generator" named iota to create a sequence of related constant variables
	const (
		sunday = iota // in this example, sunday would be assigned 0, monday would be assigned 1 etc.
		monday
		tuesday
		wednesday
		thursday
		friday
		saturdxay
	)
}
