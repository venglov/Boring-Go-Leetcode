package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func findDifferentBinaryString(nums []string) string {
	nums10 := make([]int, len(nums))
	var n64 int64
	for i, n := range nums {
		n64, _ = strconv.ParseInt(n, 2, 32)
		nums10[i] = int(n64)
	}
	sort.Ints(nums10)

	for i := 0; i < len(nums10); i++ {
		if i != nums10[i] {
			binary := strconv.FormatInt(int64(i), 2)
			zerosLen := len(nums) - len(binary)
			return strings.Repeat("0", zerosLen) + binary
		}
	}

	return strings.Repeat("1", len(nums))
}

func main() {
	nums := []string{"00", "01"}
	result := findDifferentBinaryString(nums)
	fmt.Println(result)
}
