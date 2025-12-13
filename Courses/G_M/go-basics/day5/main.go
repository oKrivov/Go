package main

import (
	"fmt"
	"go-basics/day5/mathutils"
)

type Counter struct {
	value int
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	x, y := 5, 10

	fmt.Println(x, y)
	swap(&x, &y)

	c := Counter{value: 0}
	c.Inc()
	c.Inc()
	fmt.Println(c.Get())
	c.Reset()
	fmt.Println(c.Get())

	fmt.Println(mathutils.Max(x, y))
	fmt.Println(mathutils.Min(x, y))
	fmt.Println(mathutils.Sum(nums))
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func (c *Counter) Inc() {
	c.value++
}

func (c *Counter) Get() int {
	return c.value
}

func (c *Counter) Reset() {
	c.value = 0
}
