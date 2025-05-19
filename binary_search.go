package main

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
