package _139

import "testing"

func TestWordBreak(t *testing.T) {
	if !wordBreak("catsandog", []string{"cat", "sand", "og"}) {
		t.Errorf("expect true")
	}
	if wordBreak("aaaaaaa", []string{"aaaa", "aa"}) {
		t.Errorf("expect false")
	}
	if !wordBreak("aaaaaaa", []string{"aaaa", "aaa"}) {
		t.Errorf("expect true")
	}
	if wordBreak("a", []string{"b"}) {
		t.Errorf("expect false")
	}
}
