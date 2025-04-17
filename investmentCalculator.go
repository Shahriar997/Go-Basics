package main

import (
	"fmt" // importing package
	"math"
)

func main() {
	var investmentAmount float64 = 1000
	expectedReturnRate := 5.5 // if using inferred type then do not need to use var.
	var years float64 = 10

	futureValue := calculate(
		investmentAmount,
		expectedReturnRate,
		years,
	)

	fmt.Println(futureValue)
}


func calculate(
	investmentAmount float64,
	expectedReturnRate float64,
	years float64,
) float64 {
	inflationRate := 5.5
	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)

	return futureValue / math.Pow(1 + inflationRate/100, years)
}
