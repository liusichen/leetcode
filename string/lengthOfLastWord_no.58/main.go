package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s := input.Text()
		fmt.Println(lengthOfLastWord(s))
	}
}

func lengthOfLastWord(s string) int {
	if len(s) <= 0 {
		return 0
	}
	sl := strings.Split(s, " ")
	for i := len(sl) - 1; i >= 0; i-- {
		if len(sl[i]) > 0 {
			return len(sl[i])
		}
	}
	return 0
}
