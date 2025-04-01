package main

import "fmt"

/*

	Logical operators commonly used in Go:
	&& Logical-AND
	|| Logical-OR
	 ! Logical-NOT

*/

func main() {

	var yes, no bool = true, false

	result := (yes && no)
	fmt.Println("AND LOGIC\tyes && no\t", result)

	result = (yes || no)
	fmt.Println("OR LOGIC\tyes || no\t", result)

	result = !yes
	fmt.Println("NOT Logic\tyes = ", yes, "\t!yes = ", result)

}
