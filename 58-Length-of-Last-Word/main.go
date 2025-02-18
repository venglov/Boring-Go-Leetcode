package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	n := len(s)
	count := 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if count != 0 {
				return count
			}
		} else {
			count++
		}
	}
	return count
}

func main() {
	// Test cases
	fmt.Println(lengthOfLastWord("a")) // Expected output: 5
}
