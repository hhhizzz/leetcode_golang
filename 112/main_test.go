package _112

import (
    "fmt"
    "testing"
)

func TestHasPathSum(t *testing.T) {
    root := &TreeNode{1, nil, nil}
    root.Right = &TreeNode{2, nil, nil}
    root.Right.Right = &TreeNode{3, nil, nil}
    hasSum := hasPathSum(root, 6)
    fmt.Println(hasSum)
}
