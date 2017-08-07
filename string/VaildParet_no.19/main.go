package main

import "fmt"

func main() {
	d := isValid("[])")
	fmt.Println(d)
}

func isValid(s string) bool {
	if len(s) < 0 {
		return false
	}
	if len(s) == 0 {
		return true
	}

	var symStack []byte
	symStack = append(symStack, s[0])
	for i := 1; i < len(s); i++ {
		fmt.Println(symStack)
		slen := len(symStack)
		switch s[i] {
		case ')':
			if len(symStack) == 0 {
				return false
			}
			if symStack[slen-1] == '(' && slen > 1 {
				symStack = symStack[:slen-1]
			} else if symStack[slen-1] == '(' && slen == 1 {
				symStack = []byte{}
			} else {
				return false
			}
		case '}':
			if len(symStack) == 0 {
				return false
			}
			if symStack[slen-1] == '{' && slen > 1 {
				symStack = symStack[:slen-1]
			} else if symStack[slen-1] == '{' && slen == 1 {
				symStack = []byte{}
			} else {
				return false
			}
		case ']':
			if len(symStack) == 0 {
				return false
			}
			if symStack[slen-1] == '[' && slen > 1 {
				symStack = symStack[:slen-1]
			} else if symStack[slen-1] == '[' && slen == 1 {
				symStack = []byte{}
			} else {
				return false
			}
		default:
			symStack = append(symStack, s[i])
		}
	}
	if len(symStack) > 0 {
		return false
	}
	return true
}
