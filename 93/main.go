package _93

import (
	"strconv"
)

func dfs(s string, current []string, result *[]string) {
	if len(current) == 4 && len(s) == 0 {
		ip := ""
		for i, c := range current {
			ip += c
			if i != 3 {
				ip += "."
			}
		}
		*result = append(*result, ip)
	}
	if len(current) > 4 || len(s) == 0 {
		return
	}
	for i := 1; i <= len(s) && i <= 3; i++ {
		cutString := s[:i]
		if len(cutString) > 1 && cutString[0] == '0' {
			break
		} else {
			cut, _ := strconv.Atoi(cutString)
			if cut > 255 {
				break
			}
			current = append(current, cutString)
			dfs(s[i:], current, result)
			current = current[:len(current)-1]
		}
	}
}

func restoreIpAddresses(s string) []string {
	var result []string
	dfs(s, []string{}, &result)
	return result
}
