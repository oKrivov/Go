package main

import (
	"fmt"
	"math"
	"strings"
)

// Exercise 1
type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * float64(math.Pi)
}

type Rect struct {
	Width float64
	Hight float64
}

func (r Rect) Area() float64 {
	return r.Hight * r.Width
}

// Exercise 2
type Person struct {
	Name string
}

type Employee struct {
	Person
	Position string
}

//  Exercise 3

type Logger interface {
	Log(msg string)
}

type ConsoleLogger struct{}
type UpperCaseLogger struct{}

func (ConsoleLogger) Log(msg string) {
	fmt.Println(msg)
}

func (UpperCaseLogger) Log(msg string) {
	fmt.Println(strings.ToUpper(msg))
}

func Process(l Logger, msg string) {
	l.Log(msg)
}

func main() {
	//  Exercise 1
	var s Shape

	shapes := []Shape{
		Circle{Radius: 5},
		Rect{Width: 3, Hight: 4},
	}

	for _, shape := range shapes {
		s = shape
		fmt.Println(s.Area())
	}

	//  Exercise 2
	p := Person{
		Name: "Kate",
	}

	e := Employee{
		Person:   p,
		Position: "manager",
	}

	fmt.Println(e.Name)
	fmt.Println(e.Position)

	//  Exercise 3
	Process(ConsoleLogger{}, "Console log")
	Process(UpperCaseLogger{}, "upper Console log")
}

/*
Was difficult to understaud the cocept of interface.
Mabey i need more practice is needed to consolidate the material.
And you forgot to give me practice in english and soft skills.
Could you recomend me any matireals for soft skills by myself?
*/
