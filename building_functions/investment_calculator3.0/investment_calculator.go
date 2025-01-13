package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {
	var investmentAmount, expectedReturnRate, years float64

	fmt.Print("Insert the invesment amount:")
	fmt.Scan(&investmentAmount)
	fmt.Print("Insert the Expected Return Rate:")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Insert how many years you expect to invest:")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Expected value is %.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("The future value (adjusted by inflation) is %.2f\n", futureRealValue)
	//fmt.Printf("Expected value is %.2f\n", futureValue)
	//fmt.Printf("The future value (adjusted by inflation) is %.2f\n", futureRealValue)
	print(formattedFV, formattedFRV)
}

/*func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	FV := investmentAmount * math.Pow(1.0+expectedReturnRate/100, years)
	FRV := FV / math.Pow(1+inflationRate/100, years)
	return FV, FRV
}*/
//we could do instead

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (FV float64, FRV float64) {
	FV = investmentAmount * math.Pow(1.0+expectedReturnRate/100, years)
	FRV = FV / math.Pow(1+inflationRate/100, years)
	return
}
