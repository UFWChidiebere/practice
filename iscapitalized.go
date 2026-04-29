package piscine

// Write a function IsCapitalized() that takes a string as an argument and returns true if each word in the string begins with either an uppercase letter or a non-alphabetic character.

// If any of the words begin with a lowercase letter return false.
// If the string is empty return false.
// Expected function
// func IsCapitalized(s string) bool {

// }
// Usage
// Here is a possible program to test your function:

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(IsCapitalized("Hello! How are you?"))
// 	fmt.Println(IsCapitalized("Hello How Are You"))
// 	fmt.Println(IsCapitalized("Whats 4this 100K?"))
// 	fmt.Println(IsCapitalized("Whatsthis4"))
// 	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
// 	fmt.Println(IsCapitalized(""))
// }
// And its output:

// $ go run .
// false
// true
// true
// true
// true
// false

// func IsCapitalized(s string) bool {
// 	if len(s) == 0 {
// 		return false
// 	}

// 	inWord := false

// 	for i := 0; i < len(s); i++ {
// 		c := s[i]

// 		// Detect start of a word
// 		if c != ' ' && !inWord {
// 			inWord = true

// 			// Check first character of the word
// 			if c >= 'a' && c <= 'z' {
// 				return false
// 			}
// 		}

// 		// Detect end of a word
// 		if c == ' ' {
// 			inWord = false
// 		}
// 	}

// 	return true
// }

func IsCapitalized(s string) bool {
	inWord := false

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c != ' ' && !inWord {
			if c >= 'a' && c <= 'z' {
				return false
			}
			inWord = true
		} else if c == ' ' {
			inWord = false
		}
	}
	return len(s) > 0
}
