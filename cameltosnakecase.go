package piscine

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

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(CamelToSnakeCase("HelloWorld"))
// 	fmt.Println(CamelToSnakeCase("helloWorld"))
// 	fmt.Println(CamelToSnakeCase("camelCase"))
// 	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
// 	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
// 	fmt.Println(CamelToSnakeCase("hey2"))
// }
// And its output:

// $ go run .
// Hello_World
// hello_World
// camel_Case
// CAMELtoSnackCASE
// camel_To_Snake_Case
// hey2

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	// validate camelCase rules
	for i, ch := range s {
		// only letters allowed
		if !(ch >= 'a' && ch <= 'z') && !(ch >= 'A' && ch <= 'Z') {
			return s
		}

		// must not end with uppercase
		if i == len(s)-1 && ch >= 'A' && ch <= 'Z' {
			return s
		}

		// no consecutive uppercase letters
		if i > 0 {
			prev := rune(s[i-1])
			if ch >= 'A' && ch <= 'Z' && prev >= 'A' && prev <= 'Z' {
				return s
			}
		}
	}

	result := []rune{}

	for i, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			// add underscore before uppercase (except first char)
			if i != 0 {
				result = append(result, '_')
			}
			// convert to lowercase
			result = append(result, ch+'a'-'A')
		} else {
			result = append(result, ch)
		}
	}

	return string(result)
}
