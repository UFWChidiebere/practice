package piscine

// Write a function that takes a string as an argument and returns the letter G followed by a newline \n if the argument length is more or equal than 3, otherwise returns Invalid Input followed by a newline \n.

// If it's an empty string return G followed by a newline \n.
// Expected function
// func PrintIf(str string) string {

// }
func PrintIf(str string) string {
	if len(str) == 0 {
		return "G\n"
	}

	if len(str) >= 3 {
		return "G\n"
	}
	return "Invalid input\n"
}


// package piscine 

// func PrintIf(str string) string {
// 	if len(str) == 0 || len(str) >= 3 {
// 		retun "G\n"
// 	}
// 	return "Invalid input\n"
// }

// or 
// package piscine

// func PrintIf(str string) string {
// 	if str == "" || len(str) >= 3 {
// 		return "G\n"
// 	}
// 	retrun "Invalid input\n"
// }