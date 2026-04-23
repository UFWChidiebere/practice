package main

// Write a function that converts a string from camelCase to snake_case.

// If the string is empty, return an empty string.
// If the string is not camelCase, return the string unchanged.
// If the string is camelCase, return the snake_case version of the string.
// For this exercise you need to know that camelCase has two different writing alternatives that will be accepted:

// lowerCamelCase
// UpperCamelCase
// Rules for writing in camelCase:

// The word does not end on a capitalized letter (CamelCasE).
// No two capitalized letters shall follow directly each other (CamelCAse).
// Numbers or punctuation are not allowed in the word anywhere (camelCase1).
// Expected function
// func CamelToSnakeCase(s string) string{

// }
// Usage
// Here is a possible program to test your function:

// package main

import (
	"fmt"
)

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
// And its output:

// $ go run .
// Hello_World
// hello_World
// camel_Case
// CAMELtoSnackCASE
// camel_To_Snake_Case
// hey2

func CamelToSnakeCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	if s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z' { // check if the last letter is uppercase
		return s
	}

	for i := 0; i < len(s); i++ {
		if i > 0 && s[i] >= 'A' && s[i] <= 'Z' && s[i-1] >= 'A' && s[i-1] <= 'Z' { // check if two uppercase are next to each other
			return s
		}
		if !(s[i] >= 'a' && s[i] <= 'z' || s[i] >= 'A' && s[i] <= 'Z') { // check for anything other than a letter
			return s
		}
	}

	res := ""
	for i := 0; i < len(s); i++ {
		ch := s[i]

		if ch >= 'A' && ch <= 'Z' {
			if i != 0 {
				res += "_"
			}
			res += string(ch)
		} else {
			res += string(ch)
		}
	}
	return res
}
