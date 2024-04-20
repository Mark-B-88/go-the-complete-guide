package lists

import "fmt"

func main() {
    prices := []float64{10.99, 8.99}
    prices[1] = 9.99
    prices = append(prices, 5.99, 12.99, 29.99, 100.10)
    prices = prices[1:]

    discountPrices := []float64{101.99, 80.99, 20.59}
    prices = append(prices, discountPrices...)

    fmt.Println(prices)
}