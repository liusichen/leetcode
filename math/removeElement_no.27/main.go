package main

import "fmt"

func main() {
	a := []int{3, 2, 2, 3, 4, 7, 3}
	b := removeElement(a, 3)
	fmt.Println(b)
	c := []int{1, 2}
	d := removeElement(c, 2)
	fmt.Println(d)
}

func removeElement(nums []int, val int) int {
	if len(nums) == 1 && nums[0] == val {
		return 0
	}

	j := len(nums) - 1
	for i := 0; i < len(nums) && i <= j; i++ {
		if nums[i] == val {
			for j > i && nums[j] == val {
				j--
			}
			fmt.Printf("i=%d,j=%d\n", i, j)
			if i == j && i > 0 {
				return j
			} else if i == j && i == 0 {
				return 0
			}
			nums[i] = nums[j]
			j--
		}
	}
	fmt.Println(nums[:j+1])

	return j + 1
}
