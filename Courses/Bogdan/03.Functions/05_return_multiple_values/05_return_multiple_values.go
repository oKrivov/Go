package main

import "fmt"

func processPayment(orderTotal float64, tip float64, amounPaid float64) (float64, float64) {
	totalAmauntDue := orderTotal + tip
	change := amounPaid - totalAmauntDue
	return totalAmauntDue, change
}

func main() {
	// 	6.50(orderTotal) + 2.00(tip) = 8.50(totalAmountDue)
	// 	10.00 (amountPaid)
	// 	10.00 (amountPaid) - 8.50 (totalAmountDue) = 1.50 (change)

	// 	We need to calc totalAmauntDue and change

	tottalCoast, change := processPayment(6.5, 3, 10)
	fmt.Printf("Total coast (with tip): $%.2f\nChange returned to the customer: $%.2f\n", tottalCoast, change)
}
