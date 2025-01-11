package main

import (
	"fmt"
)

func main() {
	var revenue, expenses, taxRate, ratio, earningsBeforeTax, earningsAfterTax float64

	fmt.Print("Insert your revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Insert the taxes rate paid: ")
	fmt.Scan(&taxRate)
	fmt.Print("Insert your expenses: ")
	fmt.Scan(&expenses)

	earningsBeforeTax = revenue - expenses
	fmt.Println("Your earnings before tax are: ", earningsBeforeTax)
	earningsAfterTax = revenue*(1-taxRate/100) - expenses
	fmt.Println("Your earnings after tax are: ", earningsAfterTax)
	ratio = earningsBeforeTax / earningsAfterTax
	fmt.Print("The ratio is: ", ratio)
}
