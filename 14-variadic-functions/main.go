package main

import "fmt"

func main() {
	sum := sumup(1, 10, 15, 40, -5)
	fmt.Println(sum) // 61
}

func sumup(numbers ...int) int {
	sum := 0

	for _, value := range numbers {
		sum += value // sum = sum + value
	}

	return sum
}