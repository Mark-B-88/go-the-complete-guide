# Making Sense of Recursion

Here's an example of a **_factorial function_**. We could either use a for loop for the logic, which is fine and simple, our we could also make use of **_recursion_**, where the function keeps calling itself until it reaches a **_base-case_**.

- For loop example:

```go
// Recursion, when a function calls itself

func main() {
	fact := factorial(5)
	fmt.Println(fact)
}

// factorial of 5 => 5 * 4 * 3 * 2 * 1 => 120
func factorial(number int) int {
	result := 1

	for i := 1; i <= number; i++ {
		result = result * i
	}

	return result
}
```

- Recursion example:

```go
func main() {
	fact := factorial(5)
	fmt.Println(fact)
}

func factorial(number int) int {
	// Base Case
	if number == 0 {
		return 1
	}

	// Recursion
	return number * factorial(number - 1)
}
```

---
