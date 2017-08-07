package main

import (
	"errors"
	"fmt"
)

func main() {

	a := []int{1, 3, 4, 5, 6, 7, 8}
	b := []int{11, 14, 16}
	get, err := getKthNum(a, b, 6)
	if err != nil {
		fmt.Errorf("error:%#v", err)
	}
	fmt.Printf("the num 3th is %d\n", get)
}

func getKthNum(num1, num2 []int, k int) (int, error) {
	if len(num1) == 0 {
		return num2[k-1], nil
	}
	if len(num2) == 0 {
		return num1[k-1], nil
	}
	if k == 1 {
		if num1[0] > num2[0] {
			return num2[0], nil
		}
		return num1[0], nil
	}
	if len(num1) < len(num2) {
		tmp := num1
		num1 = num2
		num2 = tmp
	}
	l1 := len(num1)
	l2 := len(num2)
	fmt.Printf("num1 len is %d\n num2 len is %d\n", l1, l2)
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Printf("k=%d\n", k)
	if k > l1+l2 {
		return 0, errors.New("k is too large")
	}
	if l1 < k/2 {
		if num2[l2-1] > num1[k-l2-1] {
			return getKthNum(num1[k-l2:], num2, l2)
		}
		return num1[k-l2-1], nil
	}
	if num2[k/2-1] > num1[k/2-1] {
		return getKthNum(num1[k/2:], num2, k-k/2)
	} else if num2[k/2-1] == num1[k/2-1] {
		return num1[k/2-1], nil
	} else {
		return getKthNum(num1, num2[k/2:], k-k/2)
	}
}
