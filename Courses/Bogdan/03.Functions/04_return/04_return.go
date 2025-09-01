package main

import "fmt"

func updateTotalPoints(currentPoints int, newPoints int) int {
	return currentPoints + newPoints
}

func calcLoyaltyPoints(amountSpend float64) int {
	loyaltyPoints := int(amountSpend * 2)
	return loyaltyPoints
}

func main() {

	totalPoints := 120
	var newlyEarnedPoints int = calcLoyaltyPoints(9.44)
	fmt.Println("Earled points today:", newlyEarnedPoints)

	totalPoints = updateTotalPoints(totalPoints, newlyEarnedPoints)
	fmt.Println("Updated loyalty points::", totalPoints)
}
