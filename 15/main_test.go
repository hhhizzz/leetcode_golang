package _15

import (
    "fmt"
    "testing"
)

func TestThreeSum(t *testing.T) {
    actual := threeSum([]int{-1, 0, 1, 2, -1, -4})
    //actual := threeSum([]int{0,0,0,0})
    fmt.Println(actual)
}
