package piscine

// Write a function called RepeatAlpha that takes a string and displays it repeating each alphabetical character as many times as its alphabetical index.

// 'a' becomes 'a', 'b' becomes 'bb', 'e' becomes 'eeeee', etc...

// Expected Function
// func RepeatAlpha(s string) string {
// }
// Usage
// Here is a possible program to test your function:

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.RepeatAlpha("abc"))
// 	fmt.Println(piscine.RepeatAlpha("Choumi."))
// 	fmt.Println(piscine.RepeatAlpha(""))
// 	fmt.Println(piscine.RepeatAlpha("abacadaba 01!"))
// }
// And its output:

// $ go run . | cat -e
// abbccc$
// CCChhhhhhhhooooooooooooooouuuuuuuuuuuuuuuuuuuuummmmmmmmmmmmmiiiiiiiii.$
// $
// abbacccaddddabba 01!$
// $

func RepeatAlpha(s string) string {
	res := ""

	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			count := int(r - 'a' + 1)
			for i := 0; i < count; i++ {
				res += string(r)
			}
		} else if r >= 'A' && r <= 'Z' {
			count := int(r - 'A' + 1)
			for i := 0; i < count; i++ {
				res += string(r)
			}
		} else {
			res += string(r)
		}
	}

	return res
}
