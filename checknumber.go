package piscine

func CheckNumber(arg string) bool {
	for _, ch := range arg {
		if ch >= '0' && ch <= '9' {
			return true
		}
	}
	return false
}
// Write a function that takes a string as an argument and returns true if the string contains any number, otherwise return false.

// Expected function
// func CheckNumber(arg string) bool {
// }




// package piscine 

// func CheckNumber(arg string) bool {
// 	for _, r := range arg {
// 		if r >= '0' && r <= '9' {
// 			return true
// 		}
// 	}
// 	return false 
// }