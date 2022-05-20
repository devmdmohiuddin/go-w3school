package main
import ("fmt")

func main() {
	var num1, num2 int = 4, 5
	fmt.Println(num1 + num2)
	fmt.Println(num1 - num2)
	fmt.Println(num1 * num2)
	fmt.Println(num1 / num2)
	fmt.Println(num1 % num2)
	
	num1++
	num1--
	fmt.Println(num1)
}