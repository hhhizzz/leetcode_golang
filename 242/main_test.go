package _242

import "testing"

func TestIsAnagram(t *testing.T) {
    if !isAnagram("anagram", "nagaram") {
        t.Errorf("wrong for anagram")
    }
    if isAnagram("rat", "car") {
        t.Errorf("wrong for rat")
    }
    if isAnagram("av", "a") {
        t.Errorf("wrong for av")
    }
}
