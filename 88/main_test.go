package _88

import "testing"

func TestMerge(t *testing.T) {
    nums1Grid := [][]int{
        {1, 2, 3, 0, 0, 0},
    }
    nums1 := []int{
        3,
    }
    nums2Grid := [][]int{
        {2, 5, 6},
    }
    num2 := []int{
        3,
    }
    expects := [][]int{
        {1, 2, 2, 3, 5, 6},
    }
    for i := range nums1Grid {
        merge(nums1Grid[i], nums1[i], nums2Grid[i], num2[i])
        if len(nums1Grid[i]) < len(expects[i]) {
            t.Errorf("expects: %v, acutal: %v", expects[i], nums1Grid[i])
        }
        for j := range nums2Grid[i] {
            if nums1Grid[i][j] != expects[i][j] {
                t.Errorf("expects: %v, acutal: %v", expects[i], nums1Grid[i])
            }
        }
    }
}
