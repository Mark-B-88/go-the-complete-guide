package main

import "fmt"

func presentOptions() {
	fmt.Println()
	fmt.Println("What do you want to do?")
	fmt.Println()
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")
	fmt.Println()
}

func closingStatement() {
	fmt.Println("Thanks for choosing our bank")
	fmt.Println()
}
