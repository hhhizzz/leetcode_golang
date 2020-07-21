package _87

import "sort"

func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	b1 := []byte(s1)
	b2 := []byte(s2)
	n := len(s1)

	sort.Slice(b1, func(i, j int) bool {
		return b1[i] < b1[j]
	})
	sort.Slice(b2, func(i, j int) bool {
		return b2[i] < b2[j]
	})
	for i := 0; i < n; i++ {
		if b1[i] != b2[i] {
			return false
		}
	}
	b1 = []byte(s1)
	b2 = []byte(s2)
	for i := 1; i < n; i++ {
		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
			return true
		}
		if isScramble(s1[:i], s2[n-i:]) && isScramble(s1[i:], s2[:n-i]) {
			return true
		}
	}
	return false
}
