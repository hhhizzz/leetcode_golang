package _189

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 2)
	fmt.Println(nums)
}
