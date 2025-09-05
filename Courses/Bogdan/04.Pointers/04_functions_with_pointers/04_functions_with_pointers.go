package main

import "fmt"

func apllyDiscount(price *float64, discountRate float64) {
	*price -= (*price * discountRate)
}

func main() {
	// 5.00
	// 10%
	// 5.00 - 5.00 * 0.10 = 5.00 -0.50 = 4.50

	var coffeePrice float64 = 5.00
	var discount float64 = 0.10
	fmt.Printf("Basic coffee price: $%.2f\n", coffeePrice)

	apllyDiscount(&coffeePrice, discount)
	fmt.Printf("Price with discount: $%.2f\n", coffeePrice)

}
