package _324

import (
	"fmt"
	"testing"
)

func TestWiggleSort(t *testing.T) {
	nums := []int{4, 5, 5, 5, 5, 6, 6, 6}
	wiggleSort(nums)
	fmt.Println(nums)
}

func TestFindNth(t *testing.T) {
	median := findNth([]int{1, 5, 2, 4, 3}, 2)
	fmt.Println(median) // 3

	median = findNth([]int{1, 4, 2, 3}, 2)
	fmt.Println(median) //3
}

func TestFindMedian(t *testing.T) {
	var median int
	median = findMedian([]int{1, 6, 2, 5, 3, 4})
	fmt.Println(median) //4
	median = findMedian([]int{1, 7, 2, 6, 3, 5, 4})
	fmt.Println(median) //4
	median = findMedian([]int{1, 3, 2, 2, 3, 1})
	fmt.Println(median) //2
}
