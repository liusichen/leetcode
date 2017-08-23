package main

import "fmt"

func main() {
	a := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(a))
}

func groupAnagrams(strs []string) [][]string {
	var tmp []string
	var res [][]string
	if len(strs) <= 0 {
		return res
	}
	tmp = append(tmp, strs[0])
	res = append(res, tmp)
	if len(strs) == 1 {
		return res
	}
	inStr := false
	for i := 1; i < len(strs); i++ {
		for key, strSet := range res {
			if isAnagrams(strSet[0], strs[i]) {
				res[key] = append(res[key], strs[i])
				fmt.Printf("strSet=%v\n", res[key])
				inStr = true
				break
			}
		}
		if !inStr {
			res = append(res, []string{strs[i]})
		}
		inStr = false
		fmt.Printf("res=%v\n", res)
	}
	return res
}

func isAnagrams(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	isList := make(map[byte]int, len(a))
	for i := 0; i < len(a); i++ {
		if _, ok := isList[a[i]]; !ok {
			isList[a[i]] = 1
			continue
		}
		isList[a[i]]++
	}
	for j := 0; j < len(b); j++ {
		if _, ok := isList[b[j]]; !ok || isList[b[j]] == 0 {
			return false
		}
		isList[b[j]]--
	}
	return true
}
