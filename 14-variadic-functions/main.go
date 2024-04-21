package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	sum := sumup(1, 10, 15, 40, -5)
	anotherSum := sumup(numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func sumup(numbers ...int) int {
	sum := 0

	for _, value := range numbers {
		sum += value
	}

	return sum
}