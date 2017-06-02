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

import "fmt"

func main() {

	//set the ugly numbers array
	var uglyNumArray [2000]int

	// we set indices for each corresponding prime factors to determine all possible ugly numbers for each prime factor
	index2, index5, index11, index17, index23 := 0, 0, 0, 0, 0
	factor2 := 2
	factor5 := 5
	factor11 := 11
	factor17 := 17
	factor23 := 23
	next_ugly_num := 1

	//declare first ugly number as 1 since it is by convention a ugly number
	uglyNumArray[0] = 1

	//iterate through the 2000 ugly numbers
	for i := 1; i < 2000; i++ {

		//find the min between all the factor subsets
		next_ugly_num = Min(factor2,
			Min(factor5,
				Min(factor11,
					Min(factor17, factor23))))

		//set the next ugly number in the array
		uglyNumArray[i] = next_ugly_num

		//if other factor subsets contain the same ugly number, we multiply it by its factor to get the next ugly number
		if factor2 == next_ugly_num {
			index2++
			factor2 = uglyNumArray[index2] * 2
		}

		if factor5 == next_ugly_num {
			index5++
			factor5 = uglyNumArray[index5] * 5
		}

		if factor11 == next_ugly_num {
			index11++
			factor11 = uglyNumArray[index11] * 11
		}

		if factor17 == next_ugly_num {
			index17++
			factor17 = uglyNumArray[index17] * 17
		}

		if factor23 == next_ugly_num {
			index23++
			factor23 = uglyNumArray[index23] * 23
		}
	}

	//print the last number in the fixed  array
	fmt.Println("The 2000'th ugly number is ", uglyNumArray[len(uglyNumArray)-1])

}

//function to determine the mininum between 2 integers
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
