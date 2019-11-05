package _151

import (
    "fmt"
    "testing"
)

func TestReverseWords(t *testing.T) {
    words := reverseWords("  hello  world!  ")
    fmt.Println(words)
}
