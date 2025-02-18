package main

import "fmt"

func majorityElement(nums []int) int {
	candidate := 0
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}
	for k, v := range countMap {
		if v > countMap[candidate] {
			candidate = k
		}
	}
	return candidate
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println("Majority Element:", majorityElement(nums))
}
