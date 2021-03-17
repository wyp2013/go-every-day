package main

import (
	"fmt"
	"sort"
)


func maximumProduct(nums []int) int {
	sort.Ints(nums)

	n := len(nums)
	valA := nums[n-1]*nums[n-2]*nums[n-3]
	valB := nums[n-1]*nums[0]*nums[1]

	if valA > valB {
		return valA
	}

	return valB
}

func main() {
	var nums []int
	nums = append(nums, []int{4,1}...)

	maximumProduct(nums)

	fmt.Println(nums)
}
