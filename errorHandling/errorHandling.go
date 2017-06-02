package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("filename.ext")

	if err == nil {
		fmt.Println(f)
	} else {
		fmt.Println(err.Error())
	}

	myError := errors.New("My error string")
	fmt.Println(myError)

	students := map[string]bool{
		"Ann":  true,
		"Mike": true}

	// if there is an item associated with that key then "ok" will be true
	attended, ok := students["Ann"]
	if ok {
		fmt.Println(ok)
		fmt.Println("attended?", attended)
	} else {
		fmt.Println("no such student")
	}

}
