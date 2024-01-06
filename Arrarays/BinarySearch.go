package main

import (
	"fmt"
	"sort"
)

func BinarySearch(numbers []int, item int) int {
	var low = 0
	var height = len(numbers) - 1
	for low <= height {
		var mid = (low + height) / 2
		var gueses = numbers[mid]
		if gueses == item {
			return gueses
		}
		if gueses > item {
			height = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	var nums = []int{10, 2, 4, 6, 9, 23}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var res = BinarySearch(nums, 23)
	fmt.Println(res)
}
