package main

// slice is an abstraction layer that sits on top of an array, at run time, it creates
// the underlying array for you

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	colors = append(colors, "Purple")
	fmt.Println(colors)

	//remove first item in the slice, syntax will have :len(colors) by default
	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	//remove last item in the slice, syntax will have "0: by default"
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	//declaring a slice with a type such as int, string etc
	// below I set a initial size of 5 and a capacity of 5
	numbers := make([]int, 5, 5)
	numbers[0] = 1
	numbers[1] = 7
	numbers[2] = 32
	numbers[3] = 82
	numbers[4] = 45
	fmt.Println(numbers)
	fmt.Println(cap(numbers))

	//appending will increase capacity to 10
	numbers = append(numbers, 235)
	fmt.Println(numbers)
	fmt.Println(cap(numbers))

	//simple sort
	sort.Ints(numbers)
	fmt.Println(numbers)
}
