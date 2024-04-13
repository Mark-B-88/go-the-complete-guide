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
		panic("Can't continue, sorry.")
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
