package _104

import (
    "fmt"
    "testing"
)

func TestMaxDepth(t *testing.T) {
    root := &TreeNode{1, nil, nil}
    root.Left = &TreeNode{2, nil, nil}
    root.Right = &TreeNode{3, nil, nil}
    root.Right.Left = &TreeNode{4, nil, nil}
    fmt.Println(maxDepth(root))
}
