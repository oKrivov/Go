package main

import "fmt"

type ConsoleLogger struct {
}

func (cl ConsoleLogger) Log(s string) {
	fmt.Println(s)
}
