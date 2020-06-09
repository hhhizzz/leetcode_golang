package _30

import "testing"

type input struct {
	s     string
	words []string
}

func TestFindSubstring(t *testing.T) {
	grid := []input{
		{"barfoofoobarthefoobarman", []string{"bar", "foo", "the"}},
		{"barfoothefoobarman", []string{"foo", "bar"}},
		{"aaa", []string{"aa", "aa"}},
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}},
		{"a", []string{}},
	}
	expects := [][]int{
		{6, 9, 12},
		{0, 9},
		{},
		{},
		{},
	}
	for i, g := range grid {
		actual := findSubstring(g.s, g.words)
		if len(actual) != len(expects[i]) {
			t.Fatalf("expects %v, actual %v", expects[i], actual)
		}
		for j := range actual {
			if actual[j] != expects[i][j] {
				t.Fatalf("expects %v, actual %v", expects[i], actual)
			}
		}
	}
}
