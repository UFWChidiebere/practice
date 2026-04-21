package piscine

// Write a function that takes two strings's and returns the number of characters that are not included in both, without repeating characters.

// If there is no unique characters return 0.
// If both strings are empty return -1.
// Expected function
// func WeAreUnique(str1 , str2 string) int {

// }
// Usage
// Here is a possible program to test your function:

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(WeAreUnique("foo", "boo"))
// 	fmt.Println(WeAreUnique("", ""))
// 	fmt.Println(WeAreUnique("abc", "def"))
// }
// And its output:

// $ go run .
// 2
// -1
// 6

func WeAreUnique(str1, str2 string) int {
	if len(str1) == 0 && len(str2) == 0 {
		return -1
	}

	m1 := make(map[rune]bool)
	m2 := make(map[rune]bool)

	for _, c := range str1 {
		m1[c] = true
	}

	for _, c := range str2 {
		m2[c] = true
	}

	count := 0

	// characters only in str1
	for c := range m1 {
		if !m2[c] {
			count++
		}
	}

	// characters only in str2
	for c := range m2 {
		if !m1[c] {
			count++
		}
	}

	return count
}
