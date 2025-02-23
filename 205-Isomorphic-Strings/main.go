package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
	asciiArr := [128]int{}
	asciiTarget := [128]int{}
	for i := range s {
		if asciiArr[s[i]] == 0 && asciiTarget[t[i]] == 0 {
			asciiArr[s[i]] = int(t[i])
			asciiTarget[t[i]] = int(s[i])
		} else if asciiArr[s[i]] != int(t[i]) || asciiTarget[t[i]] != int(s[i]) {
			return false
		}

	}
	return true
}

func main() {
	s := "aaf"
	t := "ccj"
	result := isIsomorphic(s, t)
	fmt.Println("Are the strings isomorphic?", result)
}
