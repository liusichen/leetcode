package main

import "fmt"

func main() {
	a := []int{1, 1, 1, 1, 1, 1, 1, 1, 3, 3, 4, 5, 6}
	b := removeDuplicates(a)
	fmt.Println(b)
}

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	var i, j int
	count := 1
	for i = 0; i < len(nums)-1; i++ {
		for j = i + count; j < len(nums) && nums[i] == nums[j]; j++ {
			count++
		}
		if j == len(nums) {
			fmt.Println(nums[:i+1])
			return i + 1
		}
		nums[i+1] = nums[j]
	}
	return len(nums)
}
