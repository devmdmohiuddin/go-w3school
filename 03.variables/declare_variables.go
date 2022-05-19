/*
int- stores integers (whole numbers), such as 123 or -123
float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
string - stores text, such as "Hello World". String values are surrounded by double quotes
bool- stores values with two states: true or false
*/

package main 
import ("fmt")

// Declare Varialbes
var name string = "Mohi" // with var keyword (type is string)
num := 10 // without var (type is inferred)

// Without Initival Value
var a string // output ""
var a int // output 0
var a bool // output false

// Value Assignment After Declaration
var student string
student = "Zakaria"

// Note
var location string = "Dhaka" // this format variable we can use inside and outside functions

age := 20 // this format variable we just can use inside function


func main() {
	fmt.Println("Declare Variables")
}