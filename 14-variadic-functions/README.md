# Using Variadic Functions

Let's say I wanted to simply pass the array of numbers in the `sumup()` function directly instead of assigning the number array `[]int`, how could I make this dynamic?

```go
func main() {
	numbers := []int{1, 10, 15}
	sum := sumup(numbers)
	fmt.Println(sum) // 26
}

func sumup(numbers []int) int {
	sum := 0

	for _, value := range numbers {
		sum += value // sum = sum + value
	}

	return sum
}
```

This is where **_variadic functions_** makes it possible. In the example below, notice the parameter `...int`? This is referred to as a **_variadic parameter_**, and that means that it can accept zero or more arguments of the specified type. And that is what makes it **_dynamic_** because you call it with any number of integers and it will work as expected.

Think of it being similar to JavaScript's **_spread operator_**.

```go
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
```

---

# Splitting slices into parameter values

Using the noation **_..._** in Go allows you to unpack a slice into individual values. It's useful when you have a function that accepts a variable number of arguments **_(variadic function)_**, and you want to pass elements of a slice to it as if they were individual arguments.

Although not 100% accurate, you can kind of think of it similar to JavaScript's **_rest operator_**.

```go
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
```

---
