package main

import (
	"fmt"
)

func getHappyString(n int, k int) string {
	result := make([]string, 0, k)
	currentSolution := make([]rune, 0, n)
	for _, r := range []rune{'a', 'b', 'c'} {
		currentSolution = currentSolution[:0]
		currentSolution = append(currentSolution, r)
		backtracking(currentSolution, n, k, &result)
		if len(result) == k {
			break
		}
	}
	if k > len(result) {
		return ""
	} else {
		return result[k-1]
	}
}

func backtracking(currentSolution []rune, n int, k int, result *[]string) {
	if len(*result) == k {
		return
	} else {
		if len(currentSolution) == n {
			*result = append(*result, string(currentSolution))
			return
		}
		latestRune := currentSolution[len(currentSolution)-1]
		for _, r := range []rune{'a', 'b', 'c'} {
			if latestRune != r {
				currentSolution = append(currentSolution, r)
				backtracking(currentSolution, n, k, result)
				currentSolution = currentSolution[:len(currentSolution)-1]
			}
		}
	}
}

func main() {
	n := 33
	k := 99
	result := getHappyString(n, k)
	fmt.Println(result)
}
