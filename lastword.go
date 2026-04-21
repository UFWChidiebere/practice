package piscine
// Write a function LastWord that takes a string and returns its last word followed by a \n.

// A word is a section of string delimited by spaces or by the start/end of the string.
// Expected function
// func LastWord(s string) string{

// }
// Usage
// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Print(piscine.LastWord("this        ...       is sparta, then again, maybe    not"))
// 	fmt.Print(piscine.LastWord(" lorem,ipsum "))
// 	fmt.Print(piscine.LastWord(" "))
// }
// And its output :

// $ go run . | cat -e
// not$
// lorem,ipsum$
// $
// $
// Notions
// 01-edu/z01 

func LastWord(s string) string {
	end := len(s) - 1

	// skip trailing spaces
	for end >= 0 && s[end] == ' ' {
		end--
	}

	// if no word found
	if end < 0 {
		return "\n"
	}

	start := end

	// move backwards to find start of last word
	for start >= 0 && s[start] != ' ' {
		start--
	}

	// slice and add newline
	return s[start+1:end+1] + "\n"
}