package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	shortest := strs[0]
	for _, s := range strs {
		if len(s) < len(shortest) {
			shortest = s
		}
	}

	longest := shortest
	for _, word := range strs {
		for i := range longest {
			if longest[i] != word[i] {
				longest = longest[:i]
				break
			}
		}
	}
	return string(longest)
}

func main() {
	// Example usage
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs)) // Expected output: "fl"
}
