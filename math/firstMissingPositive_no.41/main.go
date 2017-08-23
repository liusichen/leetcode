package main

import "fmt"

func main() {
	a := []int{3, 3, 2, 1, 4, -1, -3}
	fmt.Println(firstMissingPositive(a))
}

func firstMissingPositive(nums []int) int {
	if nums == nil || len(nums) <= 0 {
		return 0
	}

	nlen := len(nums)
	for i := range nums {
		if nums[i] <= 0 || nums[i] > nlen {
			nums[i] = 0
		}
	}

	pos := 0
	for i := range nums {
		if nums[i] == 0 || nums[i] == i+1 {
			continue
		}

		for nums[i] != i+1 {
			if pos = nums[i] - 1; pos < 0 || nums[pos] == nums[i] {
				break
			}

			nums[i], nums[pos] = nums[pos], nums[i]
		}
	}

	for i := range nums {
		if nums[i] == 0 || nums[i] != i+1 {
			return i + 1
		}
	}
	return nlen + 1

}
