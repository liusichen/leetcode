package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	fmt.Println(permute(a))
}

func permute(nums []int) [][]int {
	var res [][]int
	nlen := len(nums)
	dfs(&res, 0, nil, nums, nlen)
	return res
}

func dfs(res *[][]int, count int, tmp, nums []int, numsLen int) {
	fmt.Printf("tmp=%v,nums=%v,count=%d\n", tmp, nums, count)
	if count == numsLen {
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		fmt.Printf("i=%d,len(nums)=%d\n", i, len(nums))
		dfs(res, count+1, append(tmp, nums[i]), deleteIntSlice(nums, nums[i]), numsLen)

	}
}

func deleteIntSlice(s []int, key int) []int {
	var res []int
	count := 0
	for _, v := range s {
		if v == key && count == 0 {
			count++
			continue
		}
		res = append(res, v)
	}
	return res
}
