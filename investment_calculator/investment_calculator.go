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

	fmt.Println("The expected value for the future is", futureRealValue)
	fmt.Println("Expected value was ", futureValue)

}
