package main
import ("fmt")

func main() {
	// Signed Integers
  var x int = 500
  var y int = -4500
  fmt.Printf("Type: %T, value: %v", x, x)
  fmt.Printf("Type: %T, value: %v", y, y)

	// Unsigned Integers
	var i uint = 500
  var j uint = 4500
  fmt.Printf("Type: %T, value: %v", i, i)
  fmt.Printf("Type: %T, value: %v", j, j)
}