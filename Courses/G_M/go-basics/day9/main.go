package main

import (
	"fmt"
)

func main() {
	a1 := &Account{230}
	a2 := &Account{140}

	fmt.Println(Transfer(a1, a2, 40))
}
