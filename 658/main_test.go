package _658

import (
    "fmt"
    "testing"
)

func TestFindClosestElements(t *testing.T) {
    elements := findClosestElements([]int{1, 2, 3, 5, 6, 7, 8}, 3, 5)
    fmt.Println(elements)
}
