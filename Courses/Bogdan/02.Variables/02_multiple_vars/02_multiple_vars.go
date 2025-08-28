package main

import "fmt"

func main() {
	//Declare and initialize var withexplicit type
	var coffeName string = "Espresso"

	// Type inferred
	var size = "Small"

	//Short declaration and initialization. Possible only inside functions
	price := 2.50

	fmt.Println("Medium Espresso price is $2.50")
	fmt.Println(coffeName, size, "price is $", price)
	fmt.Printf("%s %s %s%.2f", coffeName, size, "price is $", price)

}
