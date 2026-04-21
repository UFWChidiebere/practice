package main
// Write a program named hiddenp that takes two strings as arguments. The program should check if the first string s1 is hidden in the second s2. s1 is considered hidden in s2 if it is possible to find each character from s1 in s2, in the same order as they appear in s1, but not necessarily consecutively.

// If s1 is hidden in s2, the program should display 1 followed by a newline.
// If s1 is not hidden in s2, the program should display 0 followed by a newline.
// If s1 is an empty string, it is considered hidden in any string.
// If the number of arguments is different from 2, the program should display nothing.
// Usage
// $ go run . "fgex.;" "tyf34gdgf;'ektufjhgdgex.;.;rtjynur6" | cat -e
// 1$
// $ go run . "abc" "2altrb53c.sse" | cat -e
// 1$
// $ go run . "abc" "btarc" | cat -e
// 0$
// $ go run . "DD" "DABC" | cat -e
// 0$
// $ go run .
// $ 

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	// Empty s1 is always hidden
	if len(s1) == 0 {
		fmt.Println(1)
		return
	}

	i := 0 // pointer for s1

	for j := 0; j < len(s2) && i < len(s1); j++ {
		if s1[i] == s2[j] {
			i++
		}
	}

	if i == len(s1) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}