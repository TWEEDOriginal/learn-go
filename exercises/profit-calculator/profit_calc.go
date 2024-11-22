package main

import (
	"fmt"
	"os"
)

func main() {
	var revenue float64 = requestUserInput("Revenue: ")
	var expenses float64 = requestUserInput("Expenses: ")
	var taxRate float64 = requestUserInput("Tax rate: ")
	var earningsPreTax, earningsPostTax, ratio float64 = calculate(revenue, expenses, taxRate)
	writeToFile(earningsPreTax, earningsPostTax, ratio)
	// set to variable
	formattedPreTax := fmt.Sprintf("Earnings before tax: %.2f \n", earningsPreTax)
	formattedPostTax := fmt.Sprintf("Earnings after tax: %.2f \n", earningsPostTax)

	// fmt.Printf("Earnings before tax: %.2f \n", earningsPreTax)
	// fmt.Printf("Earnings after tax: %.2f \n", earningsPostTax)
	fmt.Print(formattedPreTax, formattedPostTax)
	fmt.Printf("Ratio: %.2f \n", ratio)
}

func requestUserInput(description string) float64 {
	var input float64
	fmt.Print(description)
	fmt.Scan(&input)
	if input <= 0 {
		fmt.Println("Invalid amount. Must be greater than 0.")
		// trigger exit
		panic("Can't continue")
	}
	return input
}

func calculate(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsPreTax := revenue - expenses
	earningsPostTax := earningsPreTax * (1 - taxRate/100)
	ratio := earningsPreTax / earningsPostTax
	return earningsPreTax, earningsPostTax, ratio
}

func writeToFile(earningsPreTax float64, earningsPostTax float64, ratio float64) {
	amountTxt := fmt.Sprint(earningsPreTax, "\n", earningsPostTax, "\n", ratio)
	os.WriteFile("results.txt", []byte(amountTxt), 0644)
}
