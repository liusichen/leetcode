package main

import "fmt"

func main() {
	fmt.Println(isMatch("ab", "*"))
}

func isMatch(s string, p string) bool {
	i := 0
	j := 0
	starIdx := -1
	matchIdx := 0

	for i < len(s) {
		if j < len(p) && (p[j] == '?' || p[j] == s[i]) {
			i++
			j++
			continue
		}

		if j < len(p) && p[j] == '*' {
			starIdx = j
			matchIdx = i
			j++
			continue
		}
		if starIdx != -1 {
			j = starIdx + 1
			matchIdx++
			i = matchIdx
			continue
		}
		return false
	}
	for j < len(p) && p[j] == '*' {
		j++
	}

	return j == len(p)
}
