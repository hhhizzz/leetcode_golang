package _47

import (
	"fmt"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	unique := permuteUnique([]int{2, 2, 1, 1})
	fmt.Println(unique)
}
