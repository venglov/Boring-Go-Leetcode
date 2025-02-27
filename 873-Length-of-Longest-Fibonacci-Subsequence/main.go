package main

import "fmt"

func lenLongestFibSubseq(arr []int) int {
	grid := make([][]int, len(arr))
	for i := range arr {
		grid[i] = make([]int, len(arr))
	}

	longestLen := 0
	arrMap := make(map[int]int, len(arr))

	for i, n := range arr {
		arrMap[n] = i
	}

	for j := len(arr) - 1; j > 1; j-- {
		for i := len(arr) - 2; i > 0; i-- {
			if grid[i][j] == 0 {
				grid[i][j] = 2
			}

			prevNum := arr[j] - arr[i]
			prevNumI, ok := arrMap[prevNum]

			if !ok || prevNumI >= i {
				continue
			}

			grid[prevNumI][i] = grid[i][j] + 1
			if grid[prevNumI][i] > longestLen {
				longestLen = grid[prevNumI][i]
			}
		}
	}

	return longestLen
}

func main() {
	arr := []int{2, 4, 7, 8, 9, 10, 14, 15, 18, 23, 32, 50}
	result := lenLongestFibSubseq(arr)
	fmt.Println("Length of Longest Fibonacci Subsequence:", result)
}
