package main

// Write a program that takes a positive int and displays its prime factors, followed by a newline ('\n').

// Factors must be displayed in ascending order and separated by *.

// If the number of arguments is different from 1, if the argument is invalid, or if the integer does not have a prime factor, the program displays nothing.

// Usage
// $ go run . 225225
// 3*3*5*5*7*11*13
// $ go run . 8333325
// 3*3*5*5*7*11*13*37
// $ go run . 9539
// 9539
// $ go run . 804577
// 804577
// $ go run . 42
// 2*3*7
// $ go run . a
// $ go run . 0
// $ go run . 1
// $

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 1 {
		return
	}

	first := true

	// Factorization
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			if !first {
				fmt.Print("*")
			}
			fmt.Print(i)
			first = false
			n /= i
		}
	}

	// If remainder is > 1, it's a prime
	if n > 1 {
		if !first {
			fmt.Print("*")
		}
		fmt.Print(n)
	}

	fmt.Println()
}
