package _63

import "testing"

type mn struct {
	M int
	N int
}

func TestUniquePaths(t *testing.T) {
	grids := [][][]int{
		{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		},
		{
			{1, 0},
		},
		{
			{1},
			{0},
		},
	}
	expects := []int{
		2,
		0,
		0,
	}
	for index, grid := range grids {
		actual := uniquePathsWithObstacles(grid)
		expect := expects[index]
		if expect != actual {
			t.Fatalf("tests %v ,expect %v, actual %v", grid, expect, actual)
		}
	}
}
