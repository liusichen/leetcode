package main

import "fmt"

func main() {
	x := 1223
	fmt.Println(reverse(x))
}
func reverse(x int) int {
	var res int
	for x != 0 {
		res = res*10 + x%10
		x /= 10
		if res > 2147483647 || res < -2147483648 {
			res = 0
		}
	}

	return res
}
