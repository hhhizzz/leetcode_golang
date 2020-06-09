package _239

import "testing"

type example struct {
	nums   []int
	k      int
	output []int
}

func TestMaxSlidingWindow(t *testing.T) {
	grids := []example{
		{nums: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 3, output: []int{3, 3, 5, 5, 6, 7}},
		{nums: []int{1, -1}, k: 1, output: []int{1, -1}},
		{nums: []int{7, 2, 4}, k: 2, output: []int{7, 4}},
		{nums: []int{1, 3, 1, 2, 0, 5}, k: 3, output: []int{3, 3, 2, 5}},
		{nums: []int{-6, -10, -7, -1, -9, 9, -8, -4, 10, -5, 2, 9, 0, -7, 7, 4, -2, -10, 8, 7}, k: 7, output: []int{9, 9, 10, 10, 10, 10, 10, 10, 10, 9, 9, 9, 8, 8}},
	}
	for _, grid := range grids {
		actual := maxSlidingWindow(grid.nums, grid.k)
		if len(actual) != len(grid.output) {
			t.Errorf("actual: %v, expected: %v", actual, grid.output)
			break
		}
		for i := range actual {
			if actual[i] != grid.output[i] {
				t.Errorf("actual: %v, expected: %v", actual, grid.output)
				break
			}
		}
	}
}
