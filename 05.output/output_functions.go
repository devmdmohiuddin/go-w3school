/*
Go has three functions to output text:
Print()
Println()
Printf()
*/


package main 
import ("fmt")

func main() {
	var i, j = 10, "Hello"

	// The Print() Function
	fmt.Print(i, "\n", j)

	// The Println() Function
	fmt.Println(i)
	fmt.Println(j)

	// The Printf() Function
	fmt.Printf("%v %T\n", i, i) // %v -> value
	fmt.Printf("%v %T", j, j) // %T -> type
}