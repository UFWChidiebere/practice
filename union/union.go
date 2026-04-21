package main

// Write a program that takes two string and displays, without doubles, the characters that appear in either one of the string.

// The display will be in the same order that the characters appear on the command line and will be followed by a newline ('\n').

// If the number of arguments is different from 2, then the program displays a newline ('\n').

// Usage
// $ go run . zpadinton paqefwtdjetyiytjneytjoeyjnejeyj | cat -e
// zpadintoqefwjy$
// $
// $ go run . ddf6vewg64f gtwthgdwthdwfteewhrtag6h4ffdhsd | cat -e
// df6vewg4thras$
// $
// $ go run . rien "cette phrase ne cache rien" | cat -e
// rienct phas$
// $
// $ go run . | cat -e
// $
// $ go run . rien | cat -e
// $

import (
	"fmt"
	"os"
)

func contains(s string, c byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println()
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	seen := ""

	// process first string
	for i := 0; i < len(s1); i++ {
		if !contains(seen, s1[i]) {
			fmt.Printf("%c", s1[i])
			seen += string(s1[i])
		}
	}

	// process second string
	for i := 0; i < len(s2); i++ {
		if !contains(seen, s2[i]) {
			fmt.Printf("%c", s2[i])
			seen += string(s2[i])
		}
	}

	fmt.Println()
}
