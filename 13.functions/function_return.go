package main
import ("fmt")

// Function Return
func myFunction(x int, y int) int {
  return x + y
}

// Named Return Values
func myFunction1(x int, y int) (result int) {
  result = x + y
  return
}

// Store the Return Value in a Variable
func myFunction2(x int, y int) (result int) {
  result = x + y
  return
}

// Multiple Return Values
func myFunction3(x int, y string) (result int, txt1 string) {
  result = x + x
  txt1 = y + " World!"
  return
}


func main() {
  fmt.Println(myFunction(1, 2))
  fmt.Println(myFunction1(1, 2))
	total := myFunction2(1, 2)
	fmt.Println(myFunction3(5, "Hello"))
  fmt.Println(total)
}