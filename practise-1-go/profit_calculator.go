package main

import (
	"fmt"
)

/* Created by Luso
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
	Finally, to evaluate the ratio, it was simply the division of the two previous parameters:
		ratio = earningsBeforeTaxes/earningsAfterTaxes
*/

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
