// Write a program that takes one or more arguments and that, for each argument, puts the last letter of each word in uppercase and the rest in lowercase. It displays the result followed by a newline ('\n').

// If there are no argument, the program displays nothing.

// Usage
// $ go run . "First SMALL TesT" | cat -e
// firsT smalL tesT$
// $ go run . "SEconD Test IS a LItTLE EasIEr" "bEwaRe IT'S NoT HARd WhEN " " Go a dernier 0123456789 for the road e" | cat -e
// seconD tesT iS A littlE easieR$
// bewarE it'S noT harD wheN $
//  gO A dernieR 0123456789 foR thE roaD E$
// $ go run .
// $

package main

import (
	"fmt"
	"os"
)

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

func toUpper(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return c - 32
	}
	return c
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	for argIndex := 1; argIndex < len(os.Args); argIndex++ {
		s := []byte(os.Args[argIndex])

		// lowercase everything first
		for i := 0; i < len(s); i++ {
			s[i] = toLower(s[i])
		}

		// uppercase last letter of each word
		for i := 0; i < len(s); i++ {
			if s[i] != ' ' {
				if i+1 == len(s) || s[i+1] == ' ' {
					s[i] = toUpper(s[i])
				}
			}
		}

		fmt.Println(string(s))
	}
}
