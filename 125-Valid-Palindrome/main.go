package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	var normalized []rune

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			normalized = append(normalized, r)
		}
	}
	normalizedString := string(normalized)

	for i := 0; i < int(len(normalizedString)/2); i++ {
		if normalizedString[i] != normalizedString[len(normalizedString)-i-1] {
			return false
		}
	}

	return true
}

func main() {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
	}

	for _, tc := range testCases {
		result := isPalindrome(tc.input)
		fmt.Printf("Input: %s, Expected: %v, Result: %v\n", tc.input, tc.expected, result)
	}
}
