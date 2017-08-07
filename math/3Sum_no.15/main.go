package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	res := threeSum(a)
	fmt.Println(res)
}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	tmpMap := make(map[int]map[int]int)
	var result [][]int
	sort.Ints(nums)
	fmt.Println(nums)
	for index, value := range nums {
		_, ok := tmpMap[value]
		if ok {
			continue
		}
		tmpMap[value] = make(map[int]int)
		twoSum := -value
		twoSumNums := nums[index+1:]
		if twoSumNums[0] > twoSum || len(twoSumNums) < 2 {
			break
		}
		i := 0
		j := len(twoSumNums) - 1
		for i < j {
			if twoSumNums[i]+twoSumNums[j] < twoSum {
				i++
			} else if twoSumNums[i]+twoSumNums[j] > twoSum {
				j--
			} else {
				_, ok2 := tmpMap[value][twoSumNums[i]]
				if !ok2 {
					tmpMap[value][twoSumNums[i]] = twoSumNums[j]
				}
				i++
				j--
			}
		}
	}
	for num1, keyvalue := range tmpMap {
		for num2, num3 := range keyvalue {
			result = append(result, []int{num1, num2, num3})
		}
	}
	return result
}
