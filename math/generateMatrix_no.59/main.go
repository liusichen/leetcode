package main

import "fmt"

func main() {
	fmt.Println(generateMatrix(3))
}

func generateMatrix(n int) [][]int {
	row := n
	col := n
	res := make([][]int, row)
	for i := 0; i < row; i++ {
		res[i] = make([]int, col)
	}

	num := 1
	m, n := 0, -1
	for {
		for i := 0; i < col; i++ {
			n++
			res[m][n] = num
			num++
		}
		if row--; row <= 0 {
			break
		}
		for i := 0; i < row; i++ {
			m++
			res[m][n] = num
			num++
		}
		if col--; col <= 0 {
			break
		}
		for i := 0; i < col; i++ {
			n--
			res[m][n] = num
			num++
		}
		if row--; row <= 0 {
			break
		}
		for i := 0; i < row; i++ {
			m--
			res[m][n] = num
			num++
		}
		if col--; col <= 0 {
			break
		}
	}
	return res
}
