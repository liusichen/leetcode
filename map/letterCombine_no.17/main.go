package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := letterCombinations("345")
	fmt.Printf("%v\n", a)
}

var phoneLetterMap = map[string][]string{
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	var res []string
	if len(digits) == 0 {
		return res
	}
	for i := 0; i < len(digits); i++ {
		num, _ := strconv.Atoi(string(digits[i]))
		if (num-0 > 9) || (num-0 < 2) {
			return res
		}
	}
	var tmp string
	getLetterCombinations(&res, digits, 0, tmp)
	return res

}

func getLetterCombinations(res *[]string, digits string, index int, tmp string) {
	if index == len(digits)-1 {
		for _, r := range phoneLetterMap[string(digits[index])] {
			*res = append(*res, tmp+r)
		}
		fmt.Printf("tmp is %s\n", tmp)
		return
	}
	for _, r := range phoneLetterMap[string(digits[index])] {
		getLetterCombinations(res, digits, index+1, tmp+r)
	}
	return
}
