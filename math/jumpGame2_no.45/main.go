package main

import "fmt"

func main() {
	a := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(a))
}

func jump(nums []int) int {
	if nums == nil || len(nums) <= 1 {
		return 0
	}
	routerIdx := 0
	dp := make([]int, len(nums))
	dp[0] = 0

	for i := 1; i < len(nums); i++ {
		for i > nums[routerIdx]+routerIdx {
			routerIdx++
		}
		if i < nums[routerIdx]+routerIdx {
			dp[i] = dp[routerIdx] + 1
		} else if i == nums[routerIdx]+routerIdx {
			dp[i] = dp[routerIdx] + 1
			routerIdx++
		}
	}
	return dp[len(nums)-1]
}
