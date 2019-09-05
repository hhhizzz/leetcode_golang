package _123

import "testing"

func TestMaxProfit(t *testing.T) {
    inputGrid := [][]int{
        {3, 3, 5, 0, 0, 3, 1, 4},
        {1, 2, 3, 4, 5},
        {7, 6, 4, 3, 1},
    }
    expectedGrid := []int{6, 4, 0}
    for i := range inputGrid {
        actual := maxProfit(inputGrid[i])
        if actual != expectedGrid[i] {
            t.Errorf("expected: %d, acutal: %d", expectedGrid[i], actual)
        }
    }
}
