package piscine

// Write a function that takes a string and return a string containing its first word, followed by a newline ('\n').

// A word is a sequence of characters delimited by spaces or by the start/end of the argument.
// Expected Function
// func FirstWord(s string) string {
//     // ...
// }
// Usage
// Here is a possible way to test your function:

// package main

// import (
//     "fmt"

//     "piscine"
// )

// func main() {
//     fmt.Print(piscine.FirstWord("hello there"))
//     fmt.Print(piscine.FirstWord(""))
//     fmt.Print(piscine.FirstWord("hello   .........  bye"))
// }
// And its output:

// $ go run .
// hello

// hello
// $

func FirstWord(s string) string {
	if len(s) == 0 {
		return "\n"
	}

	i := 0

	// Skip leading spaces
	for i < len(s) && s[i] == ' ' {
		i++
	}

	start := i

	// Find end of first word
	for i < len(s) && s[i] != ' ' {
		i++
	}

	// If no word found
	if start == i {
		return "\n"
	}

	return s[start:i] + "\n"
}

// func LastWord(s string) string {
// 	if len(s) == 0 {
// 		return "\n"
// 	}

// 	i := len(s) - 1

// 	// Skip trailing spaces
// 	for i >= 0 && s[i] == ' ' {
// 		i--
// 	}

// 	end := i

// 	// Find start of last word
// 	for i >= 0 && s[i] != ' ' {
// 		i--
// 	}

// 	start := i + 1

// 	// If no word found
// 	if start > end {
// 		return "\n"
// 	}

// 	return s[start:end+1] + "\n"
// }