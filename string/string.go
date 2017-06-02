package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "implicitly typed string"
	fmt.Printf("str1: %v:%T\n", str1, str1)

	str2 := "explicitly typed string"
	fmt.Printf("str2: %v:%T\n", str2, str2)

	fmt.Println(strings.ToUpper(str1))
	fmt.Println(strings.Title(str1))

	lValue := "hello"
	uValue := "HELLO"
	fmt.Println("Equal?", (lValue == uValue))
	fmt.Println("Equal Non-Case-Sensitive?", strings.EqualFold(lValue, uValue))
}
