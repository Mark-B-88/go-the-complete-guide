package main

import "fmt"

func main() {
	// result := add(1,2)
	// result += 1

	// result := add("Hello ", "World ")

	result := add(54.5, 63.9)

	fmt.Println(result)
}

func add[T int|float64|string](a, b T) T {
	return a + b
}