package main

import "fmt"

func main() {
	fmt.Println(intToRoman(1234))
}

func intToRoman(num int) string {
	thousand := num / 1000
	hundred := num % 1000 / 100
	ten := num % 100 / 10
	single := num % 10

	thousandStr := []string{"M", "MM", "MMM"}
	hundredStr := []string{"C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tenStr := []string{"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	singleStr := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	var ret string
	if thousand > 0 {
		ret += thousandStr[thousand-1]
	}
	if hundred > 0 {
		ret += hundredStr[hundred-1]
	}
	if ten > 0 {
		ret += tenStr[ten-1]
	}
	if single > 0 {
		ret += singleStr[single-1]
	}

	return ret
}
