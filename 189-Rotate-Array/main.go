package main

import "fmt"

// rotate rotates the array nums by k steps to the right.
func rotate(nums []int, k int) {
	n := len(nums)
	shift := k % n
	if shift > n/2 {
		sub := make([]int, n-shift)
		copy(sub, nums[shift:])
		for i := 1; i <= shift; i++ {
			nums[n-i] = nums[shift-i]
		}
		for i := 0; i < n-shift; i++ {
			nums[i] = sub[i]
		}
	} else {
		sub := make([]int, shift)
		copy(sub, nums[:shift+1])
		for i := 0; i < n-shift; i++ {
			nums[i] = nums[shift+i]
		}
		for i := 0; i < shift; i++ {
			nums[n-shift+i] = sub[i]
		}
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	fmt.Println("Original array:", nums)
	rotate(nums, k)
	fmt.Println("Rotated array:", nums)
}
