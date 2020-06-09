package _84

import (
	"fmt"
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	var result int

	result = largestRectangleArea([]int{1, 2, 2})
	fmt.Println(result)

	result = largestRectangleArea([]int{1, 5, 6, 4})
	fmt.Println(result)
}
