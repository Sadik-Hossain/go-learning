package main

// println function fmt package er
import "fmt"

// to run go prog. -> go run <filename>.go
// always main func. theke prog run kora shuru hobe
// func main() {
// 	fmt.Print("Hi, Im sadik")          //shudhu print new line create kore na,
// 	fmt.Println("sadik learning go!!") //println new line create kore

// }

// Tokens of go lang
/*
- keywords
- data types
- variables
- escape sequences
- operators
*/
//* ------------- escape\ sequence -------------------------
func main() {
	fmt.Println("A \nB")             //* \n prints in new line
	fmt.Println("Apple \t Banana")   //* \t prints tab-spaces(4 spaces)
	fmt.Println("sadik \\ .com")     //* to print "\" we need \\
	fmt.Println("\"sadik hossain\"") //* to print along with "" use (\") before and after str

}
