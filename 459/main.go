package _459

import "strings"

//暴力匹配
func repeatedSubstringPattern1(s string) bool {
	for i := 1; i <= len(s)>>1; i++ {
		if len(s)%i != 0 {
			continue
		}
		t := i
		for ; t < len(s) && s[t] == s[t%i]; t++ {

		}
		if t == len(s) {
			return true
		}
	}
	return false
}

func repeatedSubstringPattern(s string) bool {
	s2 := s + s
	s2 = s2[1 : len(s2)-1]
	if strings.Contains(s2, s) {
		return true
	}
	return false
}
