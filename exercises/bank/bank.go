package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Welcome to the bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var accountBalance, err = getBalanceToFile()

		if err != nil {
			fmt.Print("ERROR", "\n", err, "\n", "------", "\n")
		}

		var choice = requestUserChoice("Your choice: ")
		if choice == 1 {
			fmt.Println("Your Balance is", accountBalance)
		} else if choice == 2 {
			var depositAmount float64 = requestUserInput("Your deposit: ")
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			writeBalanceToFile(accountBalance)
		} else if choice == 3 {
			var withdrawAmount float64 = requestUserInput("Withdrawal amount: ")
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			writeBalanceToFile(accountBalance)
		} else {
			fmt.Println("Goodbye!")
			break
		}
	}
	fmt.Println("Thanks for choosing out bank")

}

func requestUserChoice(description string) int {
	var choice int
	fmt.Print(description)
	fmt.Scan(&choice)
	return choice
}

func requestUserInput(description string) float64 {
	var input float64
	fmt.Print(description)
	fmt.Scan(&input)
	return input
}

const accountBalanceFile = "balance.txt"

func getBalanceToFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("failed to fetch parse stored balance value")
	}
	return balance, err
}

func writeBalanceToFile(balance float64) {
	balanceTxt := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceTxt), 0644)
}
