package main

import (
	"fmt"
)

var counter int

var seen map[string]int

var char string

func SearchWord(s, t string) string {
	// Create a map to track seen target word
	seen = make(map[string]int)
	seen[char] = counter

	return "sth"
}

var pool, word = "ADOBECODEBANC", "ABC"

func main() {
	fmt.Println(SearchWord(pool, word))
}
