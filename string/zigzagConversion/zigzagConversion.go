package main

import "fmt"

func main() {
	s := "qwertyuiopasdgh"
	rs := convert(s, 5)
	fmt.Println(rs)
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	answer := make([]byte, len(s))
	current := 0
	base := (numRows - 1) * 2
	for round := 0; round < numRows; round++ {

		index := round
		for fund := (numRows - 1 - round) * 2; ; fund = base - fund {
			if fund == 0 {
				continue
			} else if index > len(s)-1 {
				break
			} else {
				answer[current] = s[index]
				current++
			}
			index += fund
		}
	}

	return string(answer)
}
