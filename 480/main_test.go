package _480

import (
	"fmt"
	"testing"
)

func TestMedianSlidingWindow(t *testing.T) {
	result := medianSlidingWindow([]int{4, 1, 7, 1, 8, 7, 8, 7, 7, 4}, 4)
	fmt.Println(result)
}
