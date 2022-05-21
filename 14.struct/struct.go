package main 
import ("fmt") 

type Person struct {
	name string
	age int
}

func main() {
	var p1 Person

	p1.name = "Mohi"
	p1.age = 28


	fmt.Println(p1.name, p1.age)
}