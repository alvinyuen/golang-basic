package main

import "fmt"

func main() {
	// var x float64 = 42
	var result string

	// you can include an initial statement as part of the if declaration.
	// Any variables will be local to the if block and the memory they use will be eligible for garbage collection
	if x := 42; x < 0 {
		result = "Less than zero"
	} else if x == 0 {
		result = "Equal zero"
	} else {
		result = "Greater than zero"
	}

	fmt.Println("Result:", result)
}
