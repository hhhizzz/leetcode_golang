package _215

import (
    "fmt"
    "testing"
)

func TestFindKthLargest(t *testing.T) {
    largest := findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 8)
    fmt.Println(largest)
}
