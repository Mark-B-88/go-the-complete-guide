# Basic Example

```go
func main() {
	var investmentAmount = 1000
	var expectedReturnRate = 5.5
	var years = 10

	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))

	fmt.Println(futureValue)
}
```

---

# Type Conversions

```go
func main() {
	var investmentAmount float64 = 1000
	var expectedReturnRate = 5.5
	var years float64 = 10

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	fmt.Println(futureValue)
}
```

---

# Using Alternative Variable Declaration Styles

```go
func main() {
	investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Println(futureValue)
}
```

---

# Working with Constants

```go
func main() {
	const inflationRate = 2.5

	var investmentAmount float64 = 1000
	years := 10.0
	expectedReturnRate := 5.5

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
```

---

# Getting User Input

`fmt.Scan()` has limitations for use of single words only. 

```go
func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Investment Amount: ") // User Message for Input
	fmt.Scan(&investmentAmount)      // Pointer

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Investment Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
```

---

# Adjusting the value of the strings

```go
func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Investment Amount: ") // User Message for Input
	fmt.Scan(&investmentAmount)      // Pointer

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Investment Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// fmt.Println("Future Value: ", futureValue) // regular string
	fmt.Printf("Future Value: %v\nFuture Value: (adjusted for inflation): %v\n", futureRealValue, futureRealValue) // formatted string
}
```

---

# Formatting the decimals

```go
fmt.Printf("Future Value: %.2f\nFuture Value: (adjusted for inflation): %.2f\n", futureRealValue, futureRealValue)
```

---

# Creating Formatted Strings

```go
formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.2f\n", futureRealValue)

fmt.Print(formattedFV, formattedRFV)
```

---

# Building Multiline Strings

```go
fmt.Printf(`
	Future Value: %.2f
	Future Value (adjusted for inflation): %.2f
`, futureValue, futureRealValue)
```

---

# Functions : Return Values & Variable Scope

```go
const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount) // Pointer

	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Investment Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.2f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) { fmt.Print(text) }

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
```

---

# An Alternative Return Value Syntax

```go
const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount) // Pointer

	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Investment Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.2f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) { fmt.Print(text) }

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return
}
```

---