package main

import (
	"fmt"
	"sort"
)

func binarySearchMain() {
	data := []int{12, 67, -1, 98, -67, 45, 89}
	sort.Ints(data)
	fmt.Println(data)
	itemToSearch := 98
	key := binarySearch(itemToSearch, data)
	fmt.Printf("%d found at position: %d", itemToSearch, key)
	fmt.Println("")
}

func binarySearch(key int, input []int) int {
	low := 0
	high := len(input) - 1
	for low <= high {
		mid := low + (high-low)/2
		if key < input[mid] {
			high = mid - 1
		} else if key > input[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
