package _300

import (
	"fmt"
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	lis := lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6})
	fmt.Println(lis)
}
