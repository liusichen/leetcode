package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := "wordgoodgoodgoodbestword"
	words := []string{
		"word",
		"good",
		"best",
		"good",
	}
	fmt.Println(findSubstring(s, words))
}

func findSubstring(s string, words []string) []int {
	var res []int
	index := 0
	if len(words) == 0 {
		return res
	}
	wordlen := len(words[0])

	count := 0
	for index < len(s) && count < 10 {
		count++
		tmpStr := string([]byte(s)[index:])
		fmt.Println(tmpStr)
		var wordIndex []int
		tmpMap := make(map[int]string)
		for _, v := range words {
			var j int
			i := strings.Index(tmpStr, v)
			if i == -1 {
				return res
			}
			_, ok := tmpMap[i]
			if ok {
				subStr := string([]byte(tmpStr)[i+1:])
				j = strings.Index(subStr, v)
				wordIndex = append(wordIndex, i+j+1)
			} else {
				tmpMap[i] = v
				wordIndex = append(wordIndex, i)
			}
		}
		sort.Ints(wordIndex)
		fmt.Println(wordIndex)
		var wordII int
		for wordII = 0; wordII < len(wordIndex)-1; wordII++ {
			if wordIndex[wordII] != wordIndex[wordII+1]-wordlen {
				break
			}
		}
		fmt.Printf("wordII=%d\n", wordII)
		if wordII == len(wordIndex)-1 {
			res = append(res, index+wordIndex[0])
			index = index + wordIndex[1]
		} else if wordII == 0 {
			index = index + wordlen
		} else {
			index = index + wordIndex[wordII]
		}
		fmt.Printf("index=%d\n", index)
	}
	return res
}
