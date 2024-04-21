package main

import "fmt"

// Recursion, when a function calls itself

func main() {
	fact := factorial(5)
	fmt.Println(fact)
}

// factorial of 5 => 5 * 4 * 3 * 2 * 1 => 120
func factorial(number int) int {
	// Base Case
	if number == 0 {
		return 1
	}

	// Recursion
	return number * factorial(number - 1)
}