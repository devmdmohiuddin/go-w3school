package main
import ("fmt")

func main() {
	var arr1 = [5]int{1, 2, 3, 4, 5}
	nameArr := [...]string{"Hamim", "Zaka", "Sumon"} // inferred length

	// Initialize Only Specific Elements
	arr2 := [5]int{1:10,2:40}

	// Access Elements of an Array
	fmt.Println(nameArr[0])

	// Change Elements of an Array
	arr1[2] = 100

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(nameArr)
}