package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 100; i += 10 {
		fmt.Println(i)
	}

	// The Range Keyword

	fruits := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}
}
