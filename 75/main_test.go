package _75

import (
    "testing"
)

func TestSortColors(t *testing.T) {
    grids := [][]int{
        {2, 0, 2, 1, 1, 0},
    }
    expects := [][]int{
        {0, 0, 1, 1, 2, 2},
    }
    for num, grid := range grids {
        sortColors(grid)
        actual := grid
        for i, e := range expects[num] {
            if e != actual[i] {
                t.Errorf("expects %v, actual: %v\n", expects, actual)
                break
            }
        }
    }
}
