package _10

import "testing"

func TestIsMatch(t *testing.T) {
	grid := [][]string{
		//{"aa", "a*"},
		//{"ab", ".*"},
		{"aaab", "c*a*b"},
	}
	for _, item := range grid {
		if !isMatch(item[0], item[1]) {
			t.Errorf("expect match: %s %s", item[0], item[1])
		}
	}

	grid2 := [][]string{
		{"mississippi", "mis*is*p*."},
	}
	for _, item := range grid2 {
		if isMatch(item[0], item[1]) {
			t.Errorf("expect not match: %s %s", item[0], item[1])
		}
	}
}
