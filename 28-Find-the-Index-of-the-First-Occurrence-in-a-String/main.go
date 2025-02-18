package main

import (
	"fmt"
)

func main() {
	// Example usage
	haystack := "gfglla"
	needle := "akf"
	result := strStr(haystack, needle)
	fmt.Println(result)
}

// strStr function to find the index of the first occurrence of needle in haystack
func strStr(haystack string, needle string) int {
	found := false
	for i := range haystack {
		for j := 0; j < len(needle); j++ {
			found = true
			if i+j == len(haystack) {
				return -1
			}
			if needle[j] != haystack[i+j] {
				found = false
				break
			}
		}
		if found {
			return i
		}
	}

	return -1
}
