# Functions as Values & Function Types

- Functions are 1st class values
- We can use functions themselves as parameter values for other functions

Let's say we wanted a function that increases numbers within another function by 2, we could do something really easy, like this:

```go
func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := doubleNumbers(&numbers)
	fmt.Println(doubled)
}

func doubleNumbers(numbers *[]int) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		dNumbers = append(dNumbers, value*2)
	}

	return dNumbers
}
```

Or we could use a **_function type_**, like this:

```go
package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}

	// Storing Values into Variables
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)
}

// Generic Function that takes in a Function Type
func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}

	return dNumbers
}

// Regular Functions
func double(number int) int { return number * 2 }
func triple(number int) int { return number * 3 }
```

But what is the point of doing this though? Notice the `triple()` function? If we don't introduce a **_function type_** of `transform func(int)`, that means that we would have to write another regular function asides from `transformNumbers()`, and that would make the code **_repetitive_**

You could also pass a custom type as well :

```go
type transformFn func(int) int

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}

	return dNumbers
}
```

---
