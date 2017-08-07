package main

import "fmt"

func main() {
	n := []int{2, 3, 4, 1, 2}
	fmt.Println(maxArea(n))
}

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	area := 0

	for left < right {
		var h int
		if height[left] < height[right] {
			h = height[left]
		} else {
			h = height[right]
		}
		tmpArea := h * (right - left)
		fmt.Println(tmpArea)
		if tmpArea > area {
			area = tmpArea
		}
		for ; left < right && (height[left] <= h); left++ {
		}
		for ; left < right && (height[right] <= h); right-- {
		}
	}
	return area
}
