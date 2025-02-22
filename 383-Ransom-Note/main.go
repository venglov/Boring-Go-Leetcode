package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	magMap := make(map[rune]int, 26)
	for _, r := range magazine {
		magMap[r] = magMap[r] + 1
	}
	for _, r := range ransomNote {
		if magMap[r] != 0 {
			magMap[r] = magMap[r] - 1
		} else {
			return false
		}
	}
	return true
}

func main() {
	ransomNote := "examplekek"
	magazine := "examplemagazine"
	result := canConstruct(ransomNote, magazine)
	fmt.Println(result)
}
