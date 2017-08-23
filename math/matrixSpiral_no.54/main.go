package main

import "fmt"

func main() {
	matrix := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println(spiralOrder(matrix))
}

func spiralOrder(matrix [][]int) []int {
	row := len(matrix)
	if row <= 1 {
		return matrix[0]
	}
	var res []int
	col := len(matrix[0])
	m, n := 0, -1
	for {
		for i := 0; i < col; i++ {
			n++
			res = append(res, matrix[m][n])
		}
		if row--; row <= 0 {
			break
		}
		for i := 0; i < row; i++ {
			m++
			res = append(res, matrix[m][n])
		}
		if col--; col <= 0 {
			break
		}
		for i := 0; i < col; i++ {
			n--
			res = append(res, matrix[m][n])
		}
		if row--; row <= 0 {
			break
		}
		for i := 0; i < row; i++ {
			m--
			res = append(res, matrix[m][n])
		}
		if col--; col <= 0 {
			break
		}
	}
	return res
}
