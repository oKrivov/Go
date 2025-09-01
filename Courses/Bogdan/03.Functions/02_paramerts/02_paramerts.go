package main

import "fmt"

func greet(coffeeShopName string) {
	var greetMessege string = "Welcome to the Coffee Shop"
	fmt.Println(greetMessege, coffeeShopName)

}

func main() {
	greet("Brew & Beans")
	greet("Latte Shop")
}
