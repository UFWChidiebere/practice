package piscine

// Write a function that simulates the behavior of the Itoa function in Go. Itoa transforms a number represented as anint in a number represented as a string.

// For this exercise the handling of the signs + or - does have to be taken into account.

// Expected function
// func Itoa(n int) string {

// }
// Usage
// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
//     fmt.Println(piscine.Itoa(12345))
//     fmt.Println(piscine.Itoa(0))
//     fmt.Println(piscine.Itoa(-1234))
//     fmt.Println(piscine.Itoa(987654321))
// }
// And its output :

// $ go run .
// 12345
// 0
// -1234
// 987654321
// $
// Notions
// strconv/Itoa


func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
	}

	result := ""

	for n > 0 {
		digit := n % 10
		result = string('0'+digit) + result
		n /= 10
	}

	if isNegative {
		result = "-" + result
	}

	return result
}
