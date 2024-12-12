package main

import (
	"bank/file_ops"
	"fmt"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	fmt.Println("Welcome to the bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	for {
		presentOptions()
		var accountBalance, err = file_ops.GetFloatFromFile(accountBalanceFile)

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
			file_ops.WriteFloatToFile(accountBalanceFile, accountBalance)
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
			file_ops.WriteFloatToFile(accountBalanceFile, accountBalance)
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
