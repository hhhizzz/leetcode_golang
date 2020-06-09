package _280

import (
	"fmt"
	"testing"
)

func TestWiggleSort(t *testing.T) {
	num := []int{3, 5, 2, 1, 6, 4}
	wiggleSort(num)
	fmt.Println(num)
}
