package main

import (
	"fmt" // importing package
	"math"
)

func main() {
	var investmentAmount float64
	expectedReturnRate := 5.5 // if using inferred type then do not need to use var.
	var years float64

	fmt.Println("Please enter your investment amount: ")
	fmt.Scan(&investmentAmount) // with & percent, we are sending the pointer or the address

	fmt.Println("Please enter for how many year you want to invest: ")
	fmt.Scan(&years)

	fmt.Println("Please enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

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
	const inflationRate = 6.5
	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)

	return futureValue / math.Pow(1 + inflationRate/100, years)
}
