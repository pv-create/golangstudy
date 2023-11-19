package main

import "fmt"

func BoubleSort(numbers []int) {
	buffer := 0
	for i := 0; i < len(numbers); i++ {
		for j := 0; j+1 < len(numbers)-i; j++ {
			if numbers[j+1] < numbers[j] {
				buffer = numbers[j+1]
				numbers[j+1] = numbers[j]
				numbers[j] = buffer
			}
		}
		fmt.Println(numbers[i])
	}
}

func main() {
	var numbers = []int{10, 2, 4, 6, 9, 23}

	BoubleSort(numbers)

	fmt.Println(numbers)
}
