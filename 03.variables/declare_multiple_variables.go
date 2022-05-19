package main

import (
	"fmt"
)

// multiple varialbe declare and initialize
var a, b, c, d int = 1, 2, 3, 4

var (
	number int
	student string
	isActive bool
)

func main() {
	var num, name = 100, "Mohi"
	num1, name1 := 200, "Hasan"
	fmt.Println(num)
	fmt.Println(num1)
	fmt.Println(name)
	fmt.Println(name1)
}
