package _78

import (
    "fmt"
    "testing"
)

func TestSubsets(t *testing.T) {
    ints := subsets([]int{1, 2, 3})
    fmt.Println(ints)
}
