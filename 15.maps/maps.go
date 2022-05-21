package main       
import ("fmt")  

func main() {
	var a = map[string]string{"brand": "ford", "mustang": "year"}

	var b = make(map[string]string) // The map is empty now
  a["brand"] = "Ford"
  a["model"] = "Mustang"
  a["year"] = "1964"

	// Iterating Over Maps
	for k, v := range a {
		fmt.Printf("%v : %v, ", k, v)
	}

	fmt.Println(a)
	fmt.Println(b)
}