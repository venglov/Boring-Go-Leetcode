package main

import (
	"fmt"
	"strconv"
	"strings"
)

func smallestNumber(pattern string) string { // Implement the function here
	var result string
	availableNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, num := range availableNumbers {
		newAvailableNumbers := make([]int, len(availableNumbers)-1)
		copy(newAvailableNumbers[:i], availableNumbers[:i])
		copy(newAvailableNumbers[i:], availableNumbers[i+1:])
		backtracking(newAvailableNumbers, []int{num}, []rune(pattern), &result)
		if result != "" {
			break
		}
	}
	return result
}

func backtracking(availableNumbers []int, currentSolution []int, restOfString []rune, result *string) {

	if len(restOfString) == 0 {
		var sb strings.Builder
		for _, num := range currentSolution {
			sb.WriteString(strconv.Itoa(num))
		}
		*result = sb.String()
	} else {
		currentRune := restOfString[0]
		lastNum := currentSolution[len(currentSolution)-1]
		newAvailableNumbers := make([]int, len(availableNumbers)-1)

		for i, num := range availableNumbers {
			if *result != "" {
				return
			}
			if num > lastNum && currentRune == 'I' {
				currentSolution = append(currentSolution, num)
				copy(newAvailableNumbers[:i], availableNumbers[:i])
				copy(newAvailableNumbers[i:], availableNumbers[i+1:])
				backtracking(newAvailableNumbers, currentSolution, restOfString[1:], result)
				currentSolution = currentSolution[:len(currentSolution)-1]
			} else if num < lastNum && currentRune == 'D' {
				currentSolution = append(currentSolution, num)
				copy(newAvailableNumbers[:i], availableNumbers[:i])
				copy(newAvailableNumbers[i:], availableNumbers[i+1:])
				backtracking(newAvailableNumbers, currentSolution, restOfString[1:], result)
				currentSolution = currentSolution[:len(currentSolution)-1]
			}
		}
	}
}

func main() {
	// Example usage
	s := "DIIDDI"
	result := smallestNumber(s)
	fmt.Println(result)
}
