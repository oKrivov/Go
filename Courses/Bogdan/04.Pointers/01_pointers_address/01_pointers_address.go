package main

import (
	"fmt"
)

func main() {
	coffee := "Espresso"
	ptr := &coffee

	fmt.Println("Coffee name:", coffee)
	fmt.Println("Memory location:", ptr)
	fmt.Printf("Poiter adress: %p\n", ptr)

	fmt.Println("------------")

	coffeeCopy := coffee
	fmt.Println("Coffee name coffeeCopy:", coffeeCopy)
	fmt.Println("Memory coffeeCopy location:", &coffeeCopy)
	fmt.Printf("Pjiter coffeeCopy adress: %p\n", &coffeeCopy)
}
