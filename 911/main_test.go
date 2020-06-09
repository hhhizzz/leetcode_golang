package _911

import (
	"fmt"
	"testing"
)

func TestQ(t *testing.T) {
	top := Constructor([]int{0, 0, 0, 0, 1}, []int{0, 6, 39, 52, 75})
	query := []int{45, 49, 59, 68, 42, 37, 99, 26, 78, 43}
	for _, q := range query {
		result := top.Q(q)
		fmt.Println(result)
	}
}
