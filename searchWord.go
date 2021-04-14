package main

import (
	"fmt"
)

func SearchWord(s, t string) string {
	// Create a map to track seen target word
	seen := make(map[string]int)
	for _, char := range t {
		key := string(char)
		seen[key] += 1
	}

	var left = 0
	var remain_chars = len(t)
	result := [2]int{0, len(s) + 1}

	// search phase
	// decrease seen and remain_char value by 1 once target char visited
	for right, char := range s {
		if _, exist := seen[string(char)]; exist {
			seen[string(char)] -= 1
			fmt.Println(seen)
			if seen[string(char)] >= 0 {
				remain_chars -= 1
				fmt.Println(remain_chars)
			}
		}
		// replace result with smaller substring
		for remain_chars == 0 {
			if (result[1] - result[0]) > (right - left) {
				result[0] = left
				result[1] = right
			}
			// backtracking phase
			if _, exist := seen[string(s[left])]; exist {
				seen[string(s[left])] += 1
				if seen[string(s[left])] > 0 {
					remain_chars += 1
				}
			}
			left += 1
		}
	} // end for

	if result[1] == len(s)+1 {
		return ""
	} else {
		return s[result[0] : result[1]+1]
	}
}

var pool, word = "ADOBECODEBANC", "ABC"

func main() {
	fmt.Println(SearchWord(pool, word))
}
