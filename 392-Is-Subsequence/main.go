package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}
	si := 0
	for _, r := range t {
		if r == rune(s[si]) {
			si++
		}
		if si == len(s) {
			return true
		}
	}
	return false
}

func main() {
	s := ""
	t := "ahbgdc"
	result := isSubsequence(s, t)
	fmt.Println(result) // Expected output: true
}
