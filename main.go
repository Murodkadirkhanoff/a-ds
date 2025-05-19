package main

import "fmt"

func main() {
	// BINARY SEARCH
	list := []int{1, 9, 13, 25, 34, 99, 1022, 1234, 1678}
	needle := 34

	if index, found := binarySearch(list, needle); found {
		fmt.Printf("Search Index: %d\n", index)
	} else {
		fmt.Println("Ma'lumot topilmadi.")
	}
	// BINARY SEARCH

	// FACTORIAL
	x := 5
	fact := factorial(x)
	fmt.Printf("Factorial of %d is: %d", x, fact)
}
