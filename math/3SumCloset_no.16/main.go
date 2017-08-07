package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 1, 0}
	target := -100
	get := threeSumClosest(nums, target)
	fmt.Println(get)
}

func threeSumClosest(nums []int, target int) int {

	res := nums[0] + nums[1] + nums[2]
	var wantsub int
	if res > target {
		wantsub = res - target
	} else {
		wantsub = target - res
	}

	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		var sub int
		begin := i + 1
		end := len(nums) - 1
		wantTwoSum := target - nums[i]

		fmt.Printf("want two sum is %d\n", wantTwoSum)
		for begin < end {
			tmp := nums[begin] + nums[end]
			if tmp > wantTwoSum {
				sub = tmp - wantTwoSum
				end--
			} else {
				sub = wantTwoSum - tmp
				begin++
			}
			fmt.Printf("the sub is %d\n", sub)
			if sub == 0 {
				return target
			}
			if sub < wantsub {
				res = nums[i] + tmp
				wantsub = sub
			}
		}
	}
	return res
}
