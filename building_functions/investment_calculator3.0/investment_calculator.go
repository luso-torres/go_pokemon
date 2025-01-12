package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount, expectedReturnRate, years float64
	const inflationRate = 6.5

	fmt.Print("Insert the invesment amount:")
	fmt.Scan(&investmentAmount)
	fmt.Print("Insert the Expected Return Rate:")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Insert how many years you expect to invest:")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow((1.0+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)

	formattedFV := fmt.Sprintf("Expected value is %.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("The future value (adjusted by inflation) is %.2f\n", futureRealValue)
	//fmt.Printf("Expected value is %.2f\n", futureValue)
	//fmt.Printf("The future value (adjusted by inflation) is %.2f\n", futureRealValue)
	print(formattedFV, formattedFRV)
	/* If you want to output a longer text on your print, instead of "" characters, utilize ``
	in this case \n is treated as a string instead of an external command. It can be useful in future
	e.g.
	fmt.Printf(`Expected value is %.2f\n
		* any text here*
	The future value (adjusted by inflation) is %.2f\n`, futureValue, futureRealValue)
	*/
}
