package main

import (
	"fmt"
)

func romanToInt(s string) int {
	symbolMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	pairMap := map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
	result := 0
	for i := 0; i < len(s); i++ {
		sub := s[i : i+2]
		if pairMap[sub] != 0 {
			result += pairMap[s[i:i+2]]
			i++
		} else {
			result += symbolMap[s[i]]
		}
	}
	return result
}

func main() {
	// Example usage
	roman := "MCMXCIV"
	result := romanToInt(roman)
	fmt.Println("The integer value of", roman, "is", result)
}
