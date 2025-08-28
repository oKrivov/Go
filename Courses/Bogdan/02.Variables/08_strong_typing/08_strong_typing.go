package main

import "fmt"

func main() {
	// Price of one cup of coffe
	price := 4.50

	//cups soldin one day
	quantity := 15

	//  total income
	total := price * float64(quantity)

	fmt.Printf("Total income during day: %.2f\n", total)
}
