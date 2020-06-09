package _54

import (
	"fmt"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	order := spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}})
	fmt.Println(order)
}
