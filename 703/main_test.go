package _703

import "testing"

func TestKth(t *testing.T) {
	arr := []int{4, 5, 8, 2}
	kthLargest := Constructor(3, arr)

	grids := []int{3, 5, 10, 9, 4}
	expects := []int{4, 5, 5, 8, 8}
	for index, grid := range grids {
		actual := kthLargest.Add(grid)
		if expects[index] != actual {
			t.Errorf("add %d, expects %d ,actual %d", grid, expects[index], actual)
		}
	}
}
