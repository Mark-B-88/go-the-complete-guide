package main

import (
	"fmt"

	"example.com/bank/fileops"
)

const accountBalanceFile string = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

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
		presentOptions()

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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		} else {
			fmt.Println()
			fmt.Println("Goodbye!")
			fmt.Println()
			break
		}
	}

	closingStatement()
}
