package _315

import "testing"

func TestCountSmaller(t *testing.T) {
    grid := [][]int{
        {5, 2, 6, 1},
    }
    expects := [][]int{
        {2, 1, 1, 0},
    }
    for i := range grid {
        actual := countSmaller(grid[i])
        for j := range expects[i] {
            if expects[i][j] != actual[j] {
                t.Errorf("expected: %v,actual: %v", expects[i], actual)
            }
        }
    }
}
