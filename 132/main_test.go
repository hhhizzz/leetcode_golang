package _132

import (
    "fmt"
    "testing"
)

func TestMinCut(t *testing.T) {
    cut := minCut("ccaacabacb") //cc | aa | cabac | b
    fmt.Println(cut)
}
