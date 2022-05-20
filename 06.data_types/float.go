package main
import ("fmt")

func main() {
	// The float32 Keyword
  var x float32 = 123.78
  var y float32 = 3.4e+38
  fmt.Printf("Type: %T, value: %v\n", x, x)
  fmt.Printf("Type: %T, value: %v", y, y)

	// The float64 Keyword
  var i float64 = 1.7e+308
  fmt.Printf("Type: %T, value: %v", i, i)
}