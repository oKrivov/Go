package main

import "fmt"

func main() {
	const rewardPoints = 10
	fmt.Printf("Default type of rewardPoints is: %T\n", rewardPoints) // int

	var totalRewardPoints float64 = 150.3

	totalRewardPoints = totalRewardPoints + rewardPoints
	fmt.Printf("Updated loyalty points: %.2f\n", totalRewardPoints)
}
