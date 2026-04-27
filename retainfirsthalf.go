package piscine

// Write a function called RetainFirstHalf() that takes a string as an argument and returns the first half of this string.

// If the length of the string is odd, round it down.
// If the string is empty, return an empty string.
// If the string length is equal to one, return the string.
// Expected function
// func RetainFirstHalf(str string) string {

// }
// Usage
// Here is a possible program to test your function:

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
// 	fmt.Println(piscine.RetainFirstHalf("A"))
// 	fmt.Println(piscine.RetainFirstHalf(""))
// 	fmt.Println(piscine.RetainFirstHalf("Hello World"))
// }
// And its output:

// $ go run . | cat -e
// This is the 1st half$
// A$
// $
// Hello$
// func RetainFirstHalf(str string) string {
// 	length := len(str)

// 	if length == 0 {
// 		return ""
// 	}

// 	if length == 1 {
// 		return str
// 	}

// 	half := length / 2
// 	return str[:half]
// }


// package piscine 

func RetainFirstHalf(str string) string {
	n := len(str)

	if n <= 1 {
		return str
	}
	return str[:n/2]
}