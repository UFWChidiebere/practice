package piscine
// Write a function that takes two integers and returns a string showing the range of numbers from the first to the second.

// The numbers must be separated by a comma and a space.
// If any of the arguments is bigger than 99 or less than 0, the function returns Invalid followed by a newline \n.
// Prepend a 0 to any number that is less than 10.
// Add a new line \n at the end of the string.
// Expected function
// func FromTo(from int, to int) string {

// }
// Usage
// Here is a possible program to test your function:

// package main

// import (
// 	"fmt"
// 	"psicine"
// )

// func main() {
// 	fmt.Print(piscine.FromTo(1, 10))
// 	fmt.Print(piscine.FromTo(10, 1))
// 	fmt.Print(piscine.FromTo(10, 10))
// 	fmt.Print(piscine.FromTo(100, 10))
// }
// And its output:

// $ go run . | cat -e
// 01, 02, 03, 04, 05, 06, 07, 08, 09, 10$
// 10, 09, 08, 07, 06, 05, 04, 03, 02, 01$
// 10$
// Invalid$

func FromTo(from int, to int) string {
	// Validate input
	if from < 0 || to < 0 || from > 99 || to > 99 {
		return "Invalid\n"
	}

	result := ""

	for i := from; i <= to; i++ {
		// Format number with leading zero if needed
		if i < 10 {
			result += "0"
		}
		result += string('0'+i/10) + string('0'+i%10)

		// Add separator if not last element
		if i != to {
			result += ", "
		}
	}

	result += "\n"
	return result
}
