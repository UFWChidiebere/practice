// Write a function that returns the first prime number that is equal or inferior to the int passed as parameter.

// If there are no primes inferior to the int passed as parameter the function should return 0.

// Expected function
// func FindPrevPrime(nb int) int {

// }
// Usage
// Here is a possible program to test your function :

// package main

// import "fmt"

// func main() {
// 	fmt.Println(FindPrevPrime(5))
// 	fmt.Println(FindPrevPrime(4))
// }
// And its output :

// $ go run .
// 5
// 3
// $

package piscine

func FindPrevPrime(nb int) int {
	if nb < 2 {
		return 0
	}

	for i := nb; i >= 2; i-- {
		if isPrime(i) {
			return i
		}
	}
	return 0
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
