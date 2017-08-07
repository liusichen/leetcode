package main

import "fmt"

func main() {
	res := []string{"11", "22"}
	sliCom(&res)
	fmt.Println(res)
}

func sliCom(res *[]string) {
	*res = append(*res, "123")
}
