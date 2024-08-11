package main

import "fmt"

func main() {

	// Slice of Int 0 through to 10
	odd_or_even := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	for value := range odd_or_even {
		if value%2 == 0 {
			fmt.Printf("%v is even \n", value)
		} else {
			fmt.Printf("%v is odd \n", value)
		}
	}
}