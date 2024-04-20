# Unpacking List Values

Here's an example of **_unpacking slices in Go_**. Notice `discountPrices...`, it has a **_notation_** which is represented by dots, similar to the JavaScript spread operator.

This is referred to in Go as either **_unpacking_** or **_expanding_** a slice.

```go
func main() {
    // Define the initial prices slice
    prices := []float64{10.99, 8.99}

    // Modify the second element of prices
    prices[1] = 9.99

    // Append additional prices to the prices slice
    prices = append(prices, 5.99, 12.99, 29.99, 100.10)

    // Remove the first element from the prices slice
    prices = prices[1:]

    // Define the discountPrices slice
    discountPrices := []float64{101.99, 80.99, 20.59}

    // Append each element of discountPrices to the prices slice
    prices = append(prices, discountPrices...)

    // Print the final prices slice
    fmt.Println(prices)
}
```
