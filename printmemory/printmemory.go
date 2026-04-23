package main

import (
	"fmt"
)

// Write a function that takes (arr [10]byte), and displays the memory as in the example.

// After displaying the memory the function must display all the ASCII graphic characters. The non printable characters must be replaced by a dot.

// The ASCII graphic characters are any characters intended to be written, printed, or otherwise displayed in a form that can be read by humans, present on the ASCII encoding.

// Expected function
// func PrintMemory(arr [10]byte) {

// }
// Usage
// Here is a possible program to test your function :

// package main

// import "piscine"

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

// And its output :

// $ go run . | cat -e
// 68 65 6c 6c$
// 6f 10 15 2a$
// 00 00$
// hello..*..$
// $

func PrintMemory(arr [10]byte) {
	// Print hex in groups of 4 per line
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%02x", arr[i])

		// space between bytes (except end of line)
		if (i+1)%4 != 0 && i != len(arr)-1 {
			fmt.Print(" ")
		}

		// newline every 4 bytes
		if (i+1)%4 == 0 {
			fmt.Println()
		}
	}

	// If last line wasn't complete, add newline
	if len(arr)%4 != 0 {
		fmt.Println()
	}

	// Print ASCII representation
	for _, b := range arr {
		if b >= 32 && b <= 126 {
			fmt.Printf("%c", b)
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}
