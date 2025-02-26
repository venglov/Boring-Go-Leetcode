package main

import (
	"fmt"
	"math"
)

func maxAbsoluteSum(nums []int) int {
	sum := 0
	cumSum := make([]int, len(nums))
	maxSum := 0
	minSum := 0
	for i, n := range nums {
		sum += n
		if sum > maxSum {
			maxSum = sum
		} else if sum < minSum {
			minSum = sum
		}
		cumSum[i] = sum
	}

	bestSubSum := maxSum - minSum
	absBestSubNegSum := int(math.Abs(float64(minSum - maxSum)))
	if absBestSubNegSum > bestSubSum {
		return absBestSubNegSum
	}
	return bestSubSum
}

func main() {
	nums := []int{0, 0}
	result := maxAbsoluteSum(nums)
	fmt.Println("Maximum Absolute Sum of Any Subarray:", result)
}
