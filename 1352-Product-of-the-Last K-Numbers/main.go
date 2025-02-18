package main

import (
	"fmt"
)

type ProductOfNumbers struct {
	numbers  []int32
	products []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{}
}

func (this *ProductOfNumbers) Add(num int) {
	this.numbers = append(this.numbers, int32(num))

	if num == 0 {
		this.numbers = []int32{}
		this.products = []int{}
		return
	}
	if len(this.products) == 0 {
		this.products = append(this.products, num)
	} else {
		this.products = append(this.products, this.products[len(this.products)-1]*num)
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	if k > len(this.products) || k > len(this.numbers) {
		return 0
	}
	if len(this.numbers) == 1 {
		return int(this.numbers[0])
	}
	if k == len(this.products) {
		return this.products[len(this.products)-1]
	}
	return this.products[len(this.products)-1] / this.products[len(this.products)-k-1]
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
func main() {
	// Example usage
	obj := Constructor()
	// Adding a mix of large numbers, small numbers, and zeros
	testNumbers := []int{1000000000, 0, 999999999, 2, 0, 3, 4, 5, 0, 6, 7, 8, 9, 10, 0, 11, 12, 13, 14, 15, 0, 16, 17, 18, 19, 20, 0, 21, 22, 23, 24, 25, 0, 26, 27, 28, 29, 30, 0, 31, 32, 33, 34, 35, 0, 36, 37, 38, 39, 40, 0, 41, 42, 43, 44, 45, 0, 46, 47, 48, 49, 50, 0, 51, 52, 53, 54, 55, 0, 56, 57, 58, 59, 60, 0, 61, 62, 63, 64, 65, 0, 66, 67, 68, 69, 70, 0, 71, 72, 73, 74, 75, 0, 76, 77, 78, 79, 80, 0, 81, 82, 83, 84, 85, 0, 86, 87, 88, 89, 90, 0, 91, 92, 93, 94, 95, 0, 96, 97, 98, 99, 100}

	for _, num := range testNumbers {
		obj.Add(num)
	}

	// Test GetProduct with various values of k
	testCases := []int{1, 2, 3, 4, 5, 10, 20, 50, 100, 150, 200}

	for _, k := range testCases {
		fmt.Printf("GetProduct(%d) = %d\n", k, obj.GetProduct(k))
	}
}
