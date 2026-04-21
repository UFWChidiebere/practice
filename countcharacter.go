package piscine

// write a function that takes a string and a character as arguments and returns the number of times the character appears in the string.

// if the character is not in the string return 0
// if the string is empty return 0
// Expected Function
// func CountChar(str string, c rune) int {
//     ...
// }
func CountChar(str string, c rune) int {
	count := 0

	for _, ch := range str {
		if ch == c {
			count ++
		}
	}
	return count

}

// package piscine

// func CountChar(str string, c rune) {
// 	count = 0
// 	for _, r := range str {
// 		if r == c {
// 			count ++
// 		}
// 	}
// 	return count
// }