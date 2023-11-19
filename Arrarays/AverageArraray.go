package main

import "fmt"

func Averege(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	return sum / len(numbers)
}

func main() {
	var numbers = []int{1, 2, 3, 4, 5, 10}

	fmt.Println(Averege(numbers))
}
