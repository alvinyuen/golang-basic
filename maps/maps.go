package main

import (
	"fmt"
	"sort"
)

func main() {
	country := make(map[string]string)
	fmt.Println(country)

	country["HK"] = "Hong Kong"
	country["JP"] = "Japan"
	country["KR"] = "Korea"

	fmt.Println(country)

	usa := country["USA"]
	fmt.Println(usa)

	delete(country, "JP")
	fmt.Println(country)

	country["SG"] = "Singapore"

	for k, v := range country {
		fmt.Printf("%v: %v\n", k, v)
	}

	keys := make([]string, len(country))
	i := 0
	for k := range country {
		keys[i] = k
		i++
	}

	sort.Strings(keys)
	fmt.Print("\nSorted\n")
	for i := range keys {
		fmt.Println(country[keys[i]])
	}
}
