package main

import "fmt"

func main() {
	ints := []int{}

	for i := 0; i < 11; i++ {
		ints = append(ints, i)
	}

	for _, i := range ints {
		isOdd := "odd"
		if i % 2 == 0 {
			isOdd = "even"
		}
		fmt.Printf("%v is %v\n", i, isOdd)
	}
}
