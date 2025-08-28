package main

import "fmt"

func main() {
	var coffeeType string = "Latte"
	var quantity int = 3
	var unitPrice float64 = 4.25

	fmt.Printf("Ordered %d %s priced$%.2f each\n", quantity, coffeeType, unitPrice)

	var (
		customerName string = "Oleg"
		tableNumber  int    = 8
		isReadyToPay        = false
	)

	fmt.Printf("Customer %s at table %d is ready to pay: %t\n", customerName, tableNumber, isReadyToPay)

	//  No unused variables compilation errorfor const
	const (
		sizeSmall  = "S"
		sizeMedium = "M"
		sizeLarge  = "L"
	)

}
