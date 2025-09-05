package main

import "fmt"

func main() {
	var coffeePrice = 4.50
	fmt.Println("Coffee price", coffeePrice)
	// Step 1
	// Compile Time (code you write):	var coffeePrice = 4.50
	// Runtime (what machine sees):		[some memory address] = 4.50

	// Step 2
	// Compile Time (code you write):	fmt.Println("Coffee price", coffeePrice)
	// Runtime (what machine sees):		fmt.Println([some mem address], [mem address (same as step 1)])
	fmt.Println("Coffee price", &coffeePrice)
	coffeePrice = 5.0
	fmt.Println("Coffee price", &coffeePrice)

	// ptrTocoffeePrice := &coffeePrice //same as next line
	var ptrTocoffeePrice *float64 = &coffeePrice

	/* go to the memory location ptrTocoffeePrice points to
	and change value in this mamory location*/
	*ptrTocoffeePrice = 7.50
	fmt.Println("Updated Coffee price value in memory", coffeePrice)
}
