package _140

import (
    "fmt"
    "testing"
)

func TestWordBreak(t *testing.T) {
    result := wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"})
    fmt.Println(result)
}
