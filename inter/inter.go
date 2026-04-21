package main
// Write a program that takes two string and displays, without doubles, the characters that appear in both string, in the order they appear in the first one.

// The display will be followed by a newline ('\n').

// If the number of arguments is different from 2, the program displays nothing.

// Usage
// $ go run . "padinton" "paqefwtdjetyiytjneytjoeyjnejeyj"
// padinto
// $ go run . ddf6vewg64f  twthgdwthdwfteewhrtag6h4ffdhsd
// df6ewg4
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
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	seen := ""

	for i := 0; i < len(s1); i++ {
		c := s1[i]

		if contains(s2, c) && !contains(seen, c) {
			fmt.Printf("%c", c)
			seen += string(c)
		}
	}

	fmt.Println()
}