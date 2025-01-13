package main

import (
	"fmt"
)

/* Created by Luso - Version 1.1
Inputs:
revenue => int/float type representing $
expenses => int/float type represeting $
taxRate => int/float in percentage values, e.g. 24 representing 24%
-------------------------------------------------------------------------------------------
Outputs:
earningsBeforeTaxes => float type represeting $
earningsAfterTaxes => float type represeting $
ratio => in percentage values, e.g. 40 representing 40%.
-------------------------------------------------------------------------------------------
Assumptions:
	The earnings before taxes were simply obtained by subtracting the variables:
		earningsBeforeTaxes = revenue - expenses
	It was assumed that the tax is discounted based on your revenues,
	not in your earnings before taxes as follows:
		earningsAfterTaxes = revenue(1-taxRate/100) - expenses
	** corrected to the formula:
		earningsAfterTaxes = EBT(1-taxRate/100)
	Finally, to evaluate the ratio, it was simply the division of the two previous parameters:
		ratio = EBT/profit
*/

func main() {
	var revenue, expenses, taxRate, earningsBeforeTax, earningsAfterTax, ratio float64
	revenue, expenses, taxRate = requestValues()

	earningsBeforeTax, earningsAfterTax, ratio = calculateProfit(revenue, expenses, taxRate)

	showValues(earningsBeforeTax, earningsAfterTax, ratio)
}

func requestValues() (revenue float64, taxRate float64, expenses float64) {
	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Tax rate: ")
	fmt.Scan(&taxRate)
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)
	return revenue, taxRate, expenses
}

func calculateProfit(revenue float64, taxRate float64, expenses float64) (EBT float64, profit float64, ratio float64) {
	EBT = revenue - expenses
	profit = revenue*(1-taxRate/100) - expenses
	ratio = EBT / profit
	return EBT, profit, ratio
}

func showValues(EBT, profit, ratio float64) {
	fmt.Printf("Your earnings before tax are: %.2f \nYour earnings after tax are: %.2f \nThe ratio is: %.2f", EBT, profit, ratio)
}
