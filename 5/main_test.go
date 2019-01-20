package _5

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	cases := []string{
		//"babad",
		//"cbbd",
		"aaabaaaa",
	}
	expects := []string{
		//"bab",
		//"bb",
		"aaabaaa",
	}
	for i, cas := range cases {
		actual := longestPalindrome(cas)
		if actual != expects[i] {
			t.Fatalf("expected %s, actual %s\n", expects[i], actual)
		}
	}
}
