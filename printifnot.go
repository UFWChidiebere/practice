package piscine

// Write a function that takes a string and returns:

// "G\n" if the string's length is less than 3 (including empty string).

// "Invalid Input\n" otherwise.

// Expected function
// func PrintIfNot(str string) string {

// }
func PrintIfNot(str string) string {
	if str == "" {
		return "G\n"
	}
	if len(str) < 3 {
		return "G\n"
	}
	return "Invalid Input\n"
}

// package piscine 

// func PrintIfNot(str string) string {
// 	if str == "" || len(str) < 3 {
// 		return "G\n"
// 	}
// 	return "Invalid Input\n"
// }

// package piscine 

// func PrintIfNot(str string) string {
// 	if len(str) == 0 || len(str) < 3 {
// 		return "G\n"
// 	}
// 	return "Invalid Input\n"
// }
