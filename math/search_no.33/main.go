package main

import "fmt"

func main() {
	a := []int{4, 7, 9, 10, 11}
	fmt.Println(search(a, 9))
	b := []int{4, 5, 6, 1, 2, 3}
	fmt.Println(search(b, 9))
	c := []int{1, 3}
	fmt.Println(search(c, 2))
}

func search(nums []int, target int) int {
	if len(nums) <= 0 || (len(nums) == 1 && target != nums[0]) {
		return -1
	}

	var pivot int
	head := 0
	tail := len(nums) - 1
	if nums[head] < nums[tail] {
		return biSearch(nums, target)
	}
	for head < tail {
		if nums[head] > nums[tail] && head == tail-1 {
			pivot = head
			break
		}
		mid := (head + tail) >> 1
		if nums[mid] > nums[head] {
			head = mid
		} else {
			tail = mid
		}
	}

	if target >= nums[0] {
		return biSearch(nums[:pivot+1], target)
	}
	return pivot + 1 + biSearch(nums[pivot+1:], target)
}

func biSearch(nums []int, target int) int {
	head := 0
	tail := len(nums) - 1
	if target > nums[tail] || target < nums[head] {
		return -1
	}
	for head <= tail {
		mid := (head + tail) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			head = mid + 1
		} else {
			tail = tail - 1
		}
	}
	return -1
}
