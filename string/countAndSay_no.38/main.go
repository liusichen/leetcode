package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countAndSay(6))
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	if n == 2 {
		return "11"
	}
	s := "11"

	var tmp string
	for i := 3; i <= n; i++ {
		fmt.Printf("the %dth count and say:\n", i)
		for j := 0; j < len(s); j++ {
			a := s[j]
			count := 0
			for j < len(s) && a == s[j] {
				count++
				j++
			}
			j--
			countStr := strconv.Itoa(count)
			tmp = tmp + countStr + string(a)
		}
		s = tmp
		tmp = ""
	}
	return s
}
