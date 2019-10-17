package _983

import (
    "fmt"
    "testing"
)

func TestMinCostTickets(t *testing.T) {
    ans := mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15})
    fmt.Println(ans)
}
