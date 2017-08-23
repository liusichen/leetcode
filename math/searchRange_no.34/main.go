package main

import "fmt"

func main() {
	a := []int{0, 1, 1, 1, 1, 1}
	fmt.Println(searchRange(a, 1))
	fmt.Println(searchRange(a, 0))
	fmt.Println(searchRange(a, 2))
}

func searchRange(nums []int, target int) []int {
	if len(nums) <= 0 {
		return []int{-1, -1}
	}
	if len(nums) == 1 && nums[0] != target {
		return []int{-1, -1}
	} else if len(nums) == 1 && nums[0] == target {
		return []int{0, 0}
	}

	find := biSearch(nums, target)
	if find == -1 {
		return []int{-1, -1}
	}
	var head, tail int
	if find > 0 {
		head = findRange(nums, 0, find-1, target, true)
	} else {
		head = 0
	}
	if find < len(nums)-1 {
		tail = findRange(nums, find+1, len(nums)-1, target, false)
	} else {
		tail = len(nums) - 1
	}
	return []int{head, tail}
}

func biSearch(nums []int, target int) int {
	head := 0
	tail := len(nums) - 1
	for head <= tail {
		mid := (head + tail) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			tail = mid - 1
		} else {
			head = mid + 1
		}
	}
	return -1
}

func findRange(nums []int, head, tail, target int, isHead bool) int {
	if nums[tail] < target && isHead {
		return tail + 1
	}
	if nums[head] == target && isHead {
		return head
	}
	if nums[head] > target && !isHead {
		return head - 1
	}
	if nums[tail] == target && !isHead {
		return tail
	}
	mid := (head + tail) >> 1
	if nums[mid] == target && isHead {
		return findRange(nums, head, mid-1, target, isHead)
	} else if nums[mid] == target && !isHead {
		return findRange(nums, mid+1, tail, target, isHead)
	}
	if nums[mid] < target {
		return findRange(nums, mid+1, tail, target, isHead)
	}
	return findRange(nums, head, mid-1, target, isHead)
}
