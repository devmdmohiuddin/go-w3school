package main
import ("fmt")

func main() {
  prices := []int{10,20,30}
  prices[2] = 50
  fmt.Println(prices[0])
  fmt.Println(prices[2])

	// Append Elements To a Slice
	myslice1 := []int{1, 2, 3, 4, 5, 6, 7}
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))

  myslice1 = append(myslice1, 20, 21)
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))

	// Append One Slice To Another Slice
	myslice3 := []int{1, 2, 3}
	myslice4 := []int{4, 5, 6}
	finalslice := append(myslice3, myslice4...)

	fmt.Printf("myslice3=%v\n", finalslice)
  fmt.Printf("length=%d\n", len(finalslice))
  fmt.Printf("capacity=%d\n", cap(finalslice))

	// Change The Length of a Slice
	arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
  slice := arr1[1:5] // Slice array
  fmt.Printf("slice = %v\n", slice)
  fmt.Printf("length = %d\n", len(slice))
  fmt.Printf("capacity = %d\n", cap(slice))

  slice = arr1[1:3] // Change length by re-slicing the array
  fmt.Printf("slice = %v\n", slice)
  fmt.Printf("length = %d\n", len(slice))
  fmt.Printf("capacity = %d\n", cap(slice))

  slice = append(slice, 20, 21, 22, 23) // Change length by appending items
  fmt.Printf("slice = %v\n", slice)
  fmt.Printf("length = %d\n", len(slice))
  fmt.Printf("capacity = %d\n", cap(slice))
}