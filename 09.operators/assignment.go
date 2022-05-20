package main
import ("fmt")

func main() {
	var x int = 4
	
	x += 3
	x -= 3
	x *= 3
	x /= 3
	x %= 3

	fmt.Println(x)
}