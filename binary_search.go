package main

import "fmt"

func main() {
	list := []int{1, 9, 13, 25, 34, 99, 1022, 1234, 1678}

	index, found := binary_search(list, 343)

	if found {
		fmt.Println(index)
	} else {
		fmt.Println("Ma'lumot topilmadi")
	}
}

func binary_search(haystack []int, needle int) (int, bool) {
	low := 0                  // pastki ko'rsatgich o dan boshlanadi
	high := len(haystack) - 1 // yuqori ko'rsatgich ro'yhatning oxirgi elementidan boshlanadi

	for low <= high { // pastki va yuqori ko'rsatgich bitta qiymatgacha kelguncha loop bajarilishi kerak
		middle := (low + high) / 2 // ro'yhatni yarmidan ajratib olamiz

		if haystack[middle] == needle {
			return middle, true // agar needle middle ga teng bo'lsa demak natijani topdik !
		}

		if needle > haystack[middle] {
			low = middle + 1 // agar needle middledan katta bo'lsa, pastki ko'rsatgichni yarmigacha suramiz, qolgani kerak emas
		} else {
			high = middle - 1 // // agar needle middledan katta bo'lsa, yuqori ko'rsatgichni yarmigacha pasaytiramiz
		}
	}
	return -1, false // ma'lumot topilmadi
}
