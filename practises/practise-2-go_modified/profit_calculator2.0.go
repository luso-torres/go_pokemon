package main

import (
	"errors"
	"fmt"
	"os"
)

/* Created by Luso - Version 2.0
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

const getFinanceFile = "finance.txt"

// func getBalanceFromFile() ([3]float64, error) {
// 	data, err := os.ReadFile(getInputsFile)

// 	if err != nil {
// 		return [3]float64{0.0, 0.0, 0.0}, errors.New("Failed to find input file.")
// 	}

// 	inputText := string(data) //[3]string{"10.0", "2.0", "3.0"}
// 	var input []float64
// 	var err1 error

// 	for _, s := range inputText {
// 		num, err := strconv.ParseFloat(string(s), 64)
// 		if err != nil {
// 			fmt.Printf("Error converting %s: %v\n", s, err)
// 			return [3]float64{0, 0, 0}, err
// 		}
// 		input = append(input, num)
// 	}

// 	if err1 != nil {
// 		return [3]float64{0.0, 0.0, 0.0}, errors.New("Failed to parse stored input values.")
// 	}
// 	var input2 [3]float64

// 	for i := range input {
// 		input2[i] = input[i]
// 	}
// 	return input2, err
// }

func writeFinanceToFile(finance *[3]float64) {
	financeText := fmt.Sprintf("The values for EBT: %.1f\nprofit: %.1f \n And ratio: %.1f %%", finance[0], finance[1], finance[2])
	os.WriteFile(getFinanceFile, []byte(financeText), 0644)
}

func main() {
	var revenue, expenses, taxRate, earningsBeforeTax, earningsAfterTax, ratio float64
	revenue, err := requestValues("Revenue: ")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	expenses, err = requestValues("Tax rate: ")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	taxRate, err = requestValues("Expenses: ")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	earningsBeforeTax, earningsAfterTax, ratio = calculateProfit(revenue, expenses, taxRate)

	showValues(earningsBeforeTax, earningsAfterTax, ratio)
	finance := [3]float64{earningsBeforeTax, earningsAfterTax, ratio}
	writeFinanceToFile(&finance)
}

func requestValues(infoText string) (float64, error) {
	fmt.Print(infoText)
	var userInput float64
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("Invalid Input. The value must be greater than zero.")
	}
	fmt.Println("The output will be the user input %f", userInput)
	return userInput, nil
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
