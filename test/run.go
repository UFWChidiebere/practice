package main

import (
	"fmt"
	piscine "practice"
)

// func main() {
// 	result := piscine.DigitLen(-100, 16)
// 	fmt.Println(result)

func main() {
	fmt.Println(piscine.DigitLen(100, 10))
	fmt.Println(piscine.DigitLen(100, 2))
	fmt.Println(piscine.DigitLen(-100, 16))
	fmt.Println(piscine.DigitLen(100, -1))
}
