package main

import "fmt"

func main() {
	n1, l1 := FullName("Alvin", "Yuen")
	fmt.Printf("FullName: %v, number of chars: %v\n", n1, l1)

	n2, l2 := FullName("Abby", "Yuen")
	fmt.Printf("FullName: %v, number of chars: %v\n", n2, l2)
}

func FullName(f, l string) (string, int) {
	full := f + " " + l
	length := len(full)
	return full, length
}

func FullNameNamedReturn(f, l string) (full string, length int) {
	full = f + " " + l
	length = len(full)
	return
}
