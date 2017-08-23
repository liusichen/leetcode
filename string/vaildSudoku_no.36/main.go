package main

func main() {

}

func isValidSudoku(board [][]byte) bool {
	if len(board) != 9 {
		return false
	}
	for i := 0; i < 9; i++ {
		rawMap := make(map[byte]int)
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			if board[i][j] < '0' || board[i][j] > '9' {
				return false
			}
			_, ok := rawMap[board[i][j]]
			if ok {
				return false
			}
			rawMap[board[i][j]] = 1
		}
	}

	for j := 0; j < 9; j++ {
		rawMap := make(map[byte]int)
		for i := 0; i < 9; i++ {
			if board[i][j] == '.' {
				continue
			}
			if board[i][j] < '0' || board[i][j] > '9' {
				return false
			}
			_, ok := rawMap[board[i][j]]
			if ok {
				return false
			}
			rawMap[board[i][j]] = 1
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			blockMap := make(map[byte]int)
			for k := i; k < i+3; k++ {
				for l := j; j < j+3; j++ {
					if board[k][l] == '.' {
						continue
					}
					_, ok := blockMap[board[k][l]]
					if ok {
						return false
					}
					blockMap[board[k][l]] = 1
				}
			}
		}
	}

	return true
}
