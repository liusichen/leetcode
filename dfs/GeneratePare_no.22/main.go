package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	var res []string
	generateOne("", &res, n, n)
	fmt.Println(res)
	return res
}

func generateOne(tmp string, res *[]string, left int, right int) {
	if left > right {
		return
	}
	if left > 0 {
		generateOne(tmp+"(", res, left-1, right)
	}
	if right > 0 {
		generateOne(tmp+")", res, left, right-1)
	}
	if left == 0 && right == 0 {
		*res = append(*res, tmp)
		return
	}
}
