package main

import "fmt"

//you define struct as its own custom type. The struct's name typically has an initial uppercase character.

type Shoe struct {
	brand string
	size  int
}

func main() {
	nike := Shoe{"Nike", 10}
	fmt.Println(nike)
	// dump the complete contents of a struct, including its field names and values with this syntax
	fmt.Printf("%+v\n", nike)
	fmt.Printf("Breed: %v\nweight: %v\n", nike.brand, nike.size)
}
