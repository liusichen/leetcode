package main

import "fmt"

func main() {
	test := make(map[int][]int)
	test[1] = []int{1, 2, 4}
	fmt.Println(test)
}
