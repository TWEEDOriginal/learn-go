package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64
	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Tax rate: ")
	fmt.Scan(&taxRate)
	earningsPreTax := revenue - expenses
	earningsPostTax := earningsPreTax * (1 - taxRate/100)
	ratio := earningsPreTax / earningsPostTax
	fmt.Printf("Earnings before tax: %.2f \n", earningsPreTax)
	fmt.Printf("Earnings after tax: %.2f \n", earningsPostTax)
	fmt.Printf("Ratio: %.2f \n", ratio)
}
