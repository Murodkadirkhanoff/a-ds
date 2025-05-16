package main

import "fmt"

func main() {
	list := []int{1, 9, 13, 25, 34, 99, 1022, 1234, 1678}
	needle := 34

	if index, found := binarySearch(list, needle); found {
		fmt.Printf("Index: %d\n", index)
	} else {
		fmt.Println("Ma'lumot topilmadi.")
	}
}

func binarySearch(haystack []int, needle int) (int, bool) {
	low, high := 0, len(haystack)-1

	for low <= high {
		mid := (low + high) / 2

		switch {
		case haystack[mid] == needle:
			return mid, true
		case needle > haystack[mid]:
			low = mid + 1
		default:
			high = mid - 1
		}
	}
	return -1, false
}
