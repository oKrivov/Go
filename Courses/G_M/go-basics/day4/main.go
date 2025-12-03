package main

import (
	"fmt"
	"strings"
)

func main() {
	sliceTask()
	mapTask()
	text := "Hello world hello Go go"
	fmt.Println(uniqueWords(text))
}

func sliceTask() {
	s := make([]int, 0, 3)
	s = append(s, 10, 20, 30, 40)
	fmt.Println("len of s:", len(s))
	fmt.Println("cap of s:", cap(s))
	fmt.Println(s)
	s = append(s[:1], s[2:]...)
	fmt.Println(s)
}

func mapTask() {
	m := make(map[string]int)
	m["apple"] = 3
	m["orange"] = 4
	m["banana"] = 2
	m["orange"] += 1
	fmt.Println(m)
	delete(m, "banana")
	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}

	if v, ok := m["banana"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("banana not found")
	}
}

func uniqueWords(text string) int {
	textLower := strings.ToLower(text)
	words := strings.Fields(textLower)
	uniqueWords := make(map[string]bool)

	for _, word := range words {
		uniqueWords[word] = true
	}

	return len(uniqueWords)
}
