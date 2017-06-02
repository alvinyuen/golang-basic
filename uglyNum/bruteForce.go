/** -------
 **The Ugly Numbers Problem**

Ugly numbers are numbers whose only prime factors are 2, 5, 11, 17 and/or 23. The sequence "`1, 2, 4, 5, 8, 10, 11, 16, 17, 20, 22, 23, 25, 32, 34, 40, 44, 46, 50, 55, ...`" shows the first 20 ugly numbers. By convention, 1 is included.

Write a simple program in Go (golang.org) to find and print the 2000’th ugly number. The output should consist of a single line as shown below, with `<number>` replaced by the number computed.

```
The 2000'th ugly number is <number>.
```

There is no input to this program.
**/

package main

import (
	"fmt"
)

func main() {
	//set the count to 1 since we know 1 is 1 by convention is already a ugly number
	uglyNumCount := 1

	//we start from number 2 to see if it is a ugly number
	currentNum := 2

	//iterate until we reach the 2000th ugly number
	for uglyNumCount < 2000 {

		if isUglyNum(currentNum) {
			//increment count when it is a ugly number
			uglyNumCount++
			//print out the 2000th ugly number
			if uglyNumCount == 2000 {
				fmt.Println("The 2000'th ugly number is ", currentNum)
			}
		}
		//increment the next number
		currentNum++
	}
}

//function to check whether it is a uglyNum by checking if it divisible by any of the prime factors
func isUglyNum(number int) bool {

	number = isDivisible(number, 2)
	number = isDivisible(number, 5)
	number = isDivisible(number, 11)
	number = isDivisible(number, 17)
	number = isDivisible(number, 23)

	if number == 1 {
		return true
	} else {
		return false
	}

}

//function to check if its divisible
func isDivisible(num int, primeFactor int) int {

	for num%primeFactor == 0 {
		num = num / primeFactor
	}

	return num
}
