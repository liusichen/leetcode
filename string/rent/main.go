package main

import "fmt"

func main() {
	fmt.Println(rent(3, 5, 100, 10))
	fmt.Println(rent(10, 3, 9, 7))
	fmt.Println(rent(3, 0, 4, 6))
	fmt.Println(rent(3, 5, 16, 2))
}

func rent(x, f, d, p int) int {
	//钱不够一天房租
	if d < x {
		return 0
	}
	//没有水果的时候，钱不够一天房租加一个水果
	if f == 0 && d < x+p {
		return 0
	}
	//水果价钱太高，买不起的
	if d < x*f+p {
		return f
	}

	return (d + f*p) / (x + p)
}
