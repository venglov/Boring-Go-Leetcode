package main

import "fmt"

func numOfSubarrays(arr []int) int {
	result := 0
	sum := 0
	oddCount := 0
	evenCount := 1

	for i := range arr {
		sum += arr[i]
		if sum%2 == 1 {
			result = (result + evenCount) % 1000000007
			oddCount++
		} else {
			result = (result + oddCount) % 1000000007
			evenCount++
		}
	}
	return result
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	result := numOfSubarrays(arr)
	fmt.Println("Number of subarrays with odd sum:", result)
}
