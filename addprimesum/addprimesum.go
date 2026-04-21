package main

// Write a program that takes a positive integer as argument and displays the sum of all prime numbers inferior or equal to it followed by a newline ('\n').

// If the number of arguments is different from 1, or if the argument is not a positive number, the program displays 0 followed by a newline.
// Usage
// $ go run . 5
// 10
// $ go run . 7
// 17
// $ go run . -2
// 0
// $ go run . 0
// 0
// $ go run .
// 0
// $ go run . 5 7
// 0
// $

import (
	"fmt"
	"os"
	"strconv"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Check argument count
	if len(os.Args) != 2 {
		fmt.Println(0)
		return
	}

	// Convert argument to integer
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Println(0)
		return
	}

	// Sum primes ≤ n
	sum := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			sum += i
		}
	}

	fmt.Println(sum)
}
