# Learning Control Structures

Here's a basic example

```go
func main() {
	var accountBalance = 1000.00

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your Choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your balance is: ", accountBalance)
	} else if choice == 2 {
		fmt.Println("Deposit Amount: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount
		fmt.Println("Balance Updated! New Amount: ", accountBalance)
	} else if choice == 3 {
		fmt.Println("Withdrawal Amount: ")
		var withDrawalAmount float64
		fmt.Scan(&withDrawalAmount)
		accountBalance -= withDrawalAmount
		fmt.Println("Balance Updated! New Amount: ", accountBalance)
	} else {
		fmt.Println("Goodbye!")
	}
}
```

---

# Nested "if" statements

```go
func main() {
	var accountBalance = 1000.00

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your Choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your balance is: ", accountBalance)
	} else if choice == 2 {
		fmt.Println("Deposit Amount: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Invalid Amount. Must be greater than 0.")
			return
		}

		accountBalance += depositAmount
		fmt.Println("Balance Updated! New Amount: ", accountBalance)
	} else if choice == 3 {
		fmt.Println("Withdrawal Amount: ")
		var withDrawalAmount float64
		fmt.Scan(&withDrawalAmount)

		if withDrawalAmount <= 0 {
			fmt.Println("Invalid Amount. Must be greater than 0.")
			return
		}

		if withDrawalAmount > accountBalance {
			fmt.Println("Invalid Amount. You can't withdraw more than you have.")
			return
		}

		accountBalance -= withDrawalAmount
		fmt.Println("Balance Updated! New Amount: ", accountBalance)
	} else {
		fmt.Println("Goodbye!")
	}
}
```

---

# Basic for loop

```go
func main() {
	var accountBalance = 1000.00

	for i := 0; i < 200; i++ {
		fmt.Println("Welcome to Go Bank!")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your balance is: ", accountBalance)
			fmt.Println()
		} else if choice == 2 {
			fmt.Println("Deposit Amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid Amount. Must be greater than 0.")
				return
			}

			accountBalance += depositAmount
			fmt.Println("Balance Updated! New Amount: ", accountBalance)
			fmt.Println()
		} else if choice == 3 {
			fmt.Println("Withdrawal Amount: ")
			var withDrawalAmount float64
			fmt.Scan(&withDrawalAmount)

			if withDrawalAmount <= 0 {
				fmt.Println("Invalid Amount. Must be greater than 0.")
				return
			}

			if withDrawalAmount > accountBalance {
				fmt.Println("Invalid Amount. You can't withdraw more than you have.")
				return
			}

			accountBalance -= withDrawalAmount
			fmt.Println("Balance Updated! New Amount: ", accountBalance)
			fmt.Println()
		} else {
			fmt.Println("Goodbye!")
			fmt.Println()
		}
	}
}
```

---

# Infinite loops, break and continue

```go
func main() {
	var accountBalance = 1000.00
	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your balance is: ", accountBalance)
			fmt.Println()
		} else if choice == 2 {
			fmt.Println("Deposit Amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid Amount. Must be greater than 0.")
				fmt.Println()
				// return
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance Updated! New Amount: ", accountBalance)
			fmt.Println()
		} else if choice == 3 {
			fmt.Println("Withdrawal Amount: ")
			var withDrawalAmount float64
			fmt.Scan(&withDrawalAmount)

			if withDrawalAmount <= 0 {
				fmt.Println("Invalid Amount. Must be greater than 0.")
				return
			}

			if withDrawalAmount > accountBalance {
				fmt.Println("Invalid Amount. You can't withdraw more than you have.")
				return
			}

			accountBalance -= withDrawalAmount
			fmt.Println("Balance Updated! New Amount: ", accountBalance)
			fmt.Println()
		} else {
			fmt.Println("Goodbye!")
			fmt.Println()
			// return // you could either return to break the loop or break as well
			break
		}
	}

	fmt.Println("Thanks for choosing our bank")
	fmt.Println()
}
```

---

# Switch Statements

```go
func main() {
	var accountBalance = 1000.00
	fmt.Println()
	fmt.Println("Welcome to Go Bank!")
	fmt.Println()

	for {
		fmt.Println("What do you want to do?")
		fmt.Println()
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")
		fmt.Println()

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println()
			fmt.Println("Your balance is: ", accountBalance)
			fmt.Println()
		case 2:
			fmt.Println("Deposit Amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println()
				fmt.Println("Invalid Amount. Must be greater than 0.")
				fmt.Println()
				continue
			}

			accountBalance += depositAmount
			fmt.Println()
			fmt.Println("Balance Updated! New Amount: ", accountBalance)
			fmt.Println()
		case 3:
			fmt.Println("Withdrawal Amount: ")
			var withDrawalAmount float64
			fmt.Scan(&withDrawalAmount)

			if withDrawalAmount <= 0 {
				fmt.Println()
				fmt.Println("Invalid Amount. Must be greater than 0.")
				fmt.Println()
				return
			}

			if withDrawalAmount > accountBalance {
				fmt.Println()
				fmt.Println("Invalid Amount. You can't withdraw more than you have.")
				fmt.Println()
				return
			}

			accountBalance -= withDrawalAmount
			fmt.Println()
			fmt.Println("Balance Updated! New Amount: ", accountBalance)
			fmt.Println()
		default:
			fmt.Println()
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			fmt.Println()
			return
		}
	}
}
```

---

# Writing to file to persist data

```go
package main

import (
	"fmt"
	"os"
)

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func main() {
	var accountBalance = 1000.00
	fmt.Println()
	fmt.Println("Welcome to Go Bank!")
	fmt.Println()

	for {
		fmt.Println()
		fmt.Println("What do you want to do?")
		fmt.Println()
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")
		fmt.Println()

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println()
			fmt.Println("Your balance is: ", accountBalance)
			fmt.Println()
		} else if choice == 2 {
			fmt.Println()
			fmt.Println("Deposit Amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid Amount. Must be greater than 0.")
				fmt.Println()
				continue
			}

			accountBalance += depositAmount
			fmt.Println()
			fmt.Println("Balance Updated! New Amount: ", accountBalance)
			fmt.Println()
			writeBalanceToFile(accountBalance)
		} else if choice == 3 {
			fmt.Println()
			fmt.Println("Withdrawal Amount: ")
			var withDrawalAmount float64
			fmt.Scan(&withDrawalAmount)

			if withDrawalAmount <= 0 {
				fmt.Println()
				fmt.Println("Invalid Amount. Must be greater than 0.")
				fmt.Println()
				return
			}

			if withDrawalAmount > accountBalance {
				fmt.Println()
				fmt.Println("Invalid Amount. You can't withdraw more than you have.")
				fmt.Println()
				return
			}

			accountBalance -= withDrawalAmount
			fmt.Println()
			fmt.Println("Balance Updated! New Amount: ", accountBalance)
			fmt.Println()
			writeBalanceToFile(accountBalance)
		} else {
			fmt.Println()
			fmt.Println("Goodbye!")
			fmt.Println()
			break
		}
	}

	fmt.Println()
	fmt.Println("Thanks for choosing our bank")
	fmt.Println()
}
```

---

# Reading from file

```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile string = "balance.txt"

func getBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)

	return balance
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance = getBalanceFromFile()
	fmt.Println()
	fmt.Println("Welcome to Go Bank!")
	fmt.Println()

	for {
		fmt.Println()
		fmt.Println("What do you want to do?")
		fmt.Println()
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")
		fmt.Println()

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println()
			fmt.Println("Your balance is: ", accountBalance)
			fmt.Println()
		} else if choice == 2 {
			fmt.Println()
			fmt.Println("Deposit Amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid Amount. Must be greater than 0.")
				fmt.Println()
				continue
			}

			accountBalance += depositAmount
			fmt.Println()
			fmt.Println("Balance Updated! New Amount: ", accountBalance)
			fmt.Println()
			writeBalanceToFile(accountBalance)
		} else if choice == 3 {
			fmt.Println()
			fmt.Println("Withdrawal Amount: ")
			var withDrawalAmount float64
			fmt.Scan(&withDrawalAmount)

			if withDrawalAmount <= 0 {
				fmt.Println()
				fmt.Println("Invalid Amount. Must be greater than 0.")
				fmt.Println()
				return
			}

			if withDrawalAmount > accountBalance {
				fmt.Println()
				fmt.Println("Invalid Amount. You can't withdraw more than you have.")
				fmt.Println()
				return
			}

			accountBalance -= withDrawalAmount
			fmt.Println()
			fmt.Println("Balance Updated! New Amount: ", accountBalance)
			fmt.Println()
			writeBalanceToFile(accountBalance)
		} else {
			fmt.Println()
			fmt.Println("Goodbye!")
			fmt.Println()
			break
		}
	}

	fmt.Println()
	fmt.Println("Thanks for choosing our bank")
	fmt.Println()
}
```

---

# Error Handling

```go
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile string = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 0, errors.New("failed to find balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 0, errors.New("failed to parse stored balance value")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println()
	}

	fmt.Println()
	fmt.Println("Welcome to Go Bank!")
	fmt.Println()

	for {
		fmt.Println()
		fmt.Println("What do you want to do?")
		fmt.Println()
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")
		fmt.Println()

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println()
			fmt.Println("Your balance is: ", accountBalance)
			fmt.Println()
		} else if choice == 2 {
			fmt.Println()
			fmt.Println("Deposit Amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid Amount. Must be greater than 0.")
				fmt.Println()
				continue
			}

			accountBalance += depositAmount
			fmt.Println()
			fmt.Println("Balance Updated! New Amount: ", accountBalance)
			fmt.Println()
			writeBalanceToFile(accountBalance)
		} else if choice == 3 {
			fmt.Println()
			fmt.Println("Withdrawal Amount: ")
			var withDrawalAmount float64
			fmt.Scan(&withDrawalAmount)

			if withDrawalAmount <= 0 {
				fmt.Println()
				fmt.Println("Invalid Amount. Must be greater than 0.")
				fmt.Println()
				return
			}

			if withDrawalAmount > accountBalance {
				fmt.Println()
				fmt.Println("Invalid Amount. You can't withdraw more than you have.")
				fmt.Println()
				return
			}

			accountBalance -= withDrawalAmount
			fmt.Println()
			fmt.Println("Balance Updated! New Amount: ", accountBalance)
			fmt.Println()
			writeBalanceToFile(accountBalance)
		} else {
			fmt.Println()
			fmt.Println("Goodbye!")
			fmt.Println()
			break
		}
	}

	fmt.Println()
	fmt.Println("Thanks for choosing our bank")
	fmt.Println()
}
```

---

# Handling traces

If you want to end the application in the event of an error that needs to be **_traced_** for debugging, simply use `panic()`

Obviously, using **_return_** will stop the application in the event of an error, but it will not provide **_error tracing_**

```go
if err != nil {
	fmt.Println("ERROR")
	fmt.Println(err)
	fmt.Println()
	// return
	panic("Can't continue, sorry.")
}
```

---

# Adding 3rd party packages

`go get github.com/Pallinder/go-randomdata`

---
