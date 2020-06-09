package algorithm

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	array := []int{9, 8, 7, 6, 5, 4, 10, 3, 2, 1}
	mergerSort(array)
	fmt.Println(array)
}
