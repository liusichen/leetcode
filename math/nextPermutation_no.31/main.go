package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 1, 3, 5, 3, 1}
	fmt.Printf("a=%v\n", a)
	nextPermutation(a)
	b := []int{1, 2, 5, 4, 3, 2, 1}
	fmt.Printf("b=%v\n", b)
	nextPermutation(b)

}

func nextPermutation(nums []int) {
	if len(nums) == 1 {
		fmt.Println(nums[0])
		return
	}

	//只要最后一个数字比倒数第二个大，那互换即可
	if nums[len(nums)-1] > nums[len(nums)-2] {
		nums[len(nums)-1], nums[len(nums)-2] = nums[len(nums)-2], nums[len(nums)-1]
		fmt.Println(nums)
		return
	}

	//寻找最后一个峰值
	a := len(nums) - 1
	for a > 0 && nums[a] <= nums[a-1] {
		a--
	}

	//如果没有下降点，即最大值
	if a == 0 {
		for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
		fmt.Println(nums)
		return
	}

	needChange1 := a - 1

	var needChange2 int
	for i := len(nums) - 1; i > needChange1; i-- {
		if nums[i] > nums[needChange1] {
			needChange2 = i
			break
		}
	}
	nums[needChange1], nums[needChange2] = nums[needChange2], nums[needChange1]
	res := nums[:needChange1+1]
	fmt.Println(res)
	tmp1 := nums[needChange1+1:]

	sort.Ints(tmp1)

	for _, v := range tmp1 {
		res = append(res, v)
	}
	fmt.Println(res)
	return

}
