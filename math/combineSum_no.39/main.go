package main

import "fmt"

func main() {
	a := []int{1, 2}
	fmt.Println(combinationSum(a, 4))
}

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var tmp []int
	dfs(&res, &tmp, candidates, target, 0)
	return res
}

func dfs(res *[][]int, tmp *[]int, candidates []int, target int, begin int) {
	fmt.Printf("res=%v\n", *res)
	fmt.Printf("tmp=%v,target=%d, candidates=%d\n", *tmp, target, candidates[begin])
	if target == 0 {
		var a []int
		for _, v := range *tmp {
			a = append(a, v)
		}
		*res = append(*res, a)
		*tmp = (*tmp)[:len(*tmp)-1]
		fmt.Println((*tmp))
		return
	}

	for i := begin; i < len(candidates); i++ {
		if target >= candidates[i] {
			*tmp = append(*tmp, candidates[i])
			dfs(res, tmp, candidates, target-candidates[i], i)
		} else {
			break
		}
	}
	if len(*tmp) > 0 {
		*tmp = (*tmp)[:len(*tmp)-1]
	}
}
