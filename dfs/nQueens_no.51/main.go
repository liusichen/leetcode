package main

import "fmt"

func main() {
	fmt.Println(solveNQueens(4))
}
func solveNQueens(n int) [][]string {
	var res [][]string
	var tmp [][]byte
	for i := 0; i < n; i++ {
		var s []byte
		for j := 0; j < n; j++ {
			s = append(s, '.')
		}
		tmp = append(tmp, s)
	}
	dfs(&res, &tmp, 0, n)
	return res
}

func dfs(res *[][]string, tmp *[][]byte, row, n int) {
	if row == n {
		var tmpStr []string
		for _, v := range *tmp {
			tmpStr = append(tmpStr, string(v))
		}
		*res = append(*res, tmpStr)
		return
	}

	for col := 0; col < n; col++ {
		if isValid(tmp, row, col, n) {
			(*tmp)[row][col] = 'Q'
			dfs(res, tmp, row+1, n)
			(*tmp)[row][col] = '.'
		}
	}
}

func isValid(tmp *[][]byte, row int, col int, n int) bool {
	for i := 0; i < row; i++ {
		if (*tmp)[i][col] == 'Q' {
			return false
		}
	}

	var i, j int
	for i, j = row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if (*tmp)[i][j] == 'Q' {
			return false
		}
	}

	for i, j = row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if (*tmp)[i][j] == 'Q' {
			return false
		}
	}
	return true
}
