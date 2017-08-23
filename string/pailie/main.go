package main

import "fmt"

func main() {
	fmt.Println(pailie("AAAA"))
	fmt.Println(pailie("ACBA"))
	fmt.Println(pailie("AABB"))
}

func pailie(s string) int {
	if len(s) <= 0 {
		return 0
	}
	blockMap := make(map[string]int)
	for i := 0; i < len(s); i++ {
		blockMap[string(s[i])]++
	}
	if len(blockMap) > 2 {
		return 0
	} else if len(blockMap) == 1 {
		return 1
	}
	return 2
}
