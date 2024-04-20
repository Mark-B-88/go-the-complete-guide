# Creating a number array

```go
package main

import "fmt"

func main() {
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)
}
```

---

# Working with arrays

This prints out an empty array with 4 slots.

```go
package main

import "fmt"

func main() {
	var productNames [4]string
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)
	fmt.Println(productNames)
}
```

In this example, even though we set 4 array slots, we don't have to use them all:

```go
func main() {
	var productNames [4]string = [4]string{"A Book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)
	fmt.Println(productNames)
}
```

You can also access a value in the array similar to other languages which are zero based:

```go
func main() {
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices[2]) // 45.99
}
```

You can also assign values to an array:

```go
func main() {
	var productNames [4]string = [4]string{"A Book"}

	productNames[2] = "A Carpet"
	fmt.Println(productNames)
}
```

## Selecting parts of arrays with slices

This is a slice method similar to Python, in this example, we are starting from the 1st item all the way to the 3rd item, in the zero based array.

```go
func main() {
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

	featuredPrices := prices[1:3] // 9.99, 45.99, 20.0
	fmt.Println(featuredPrices)
}
```

---
