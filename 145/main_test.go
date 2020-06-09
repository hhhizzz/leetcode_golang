package _145

import (
	"fmt"
	"testing"
)

func TestPostOrderTraversal(t *testing.T) {
	root := &TreeNode{1, nil, nil}
	root.Left = &TreeNode{2, nil, nil}
	root.Right = &TreeNode{3, nil, nil}
	root.Left.Right = &TreeNode{4, nil, nil}
	traversal := postorderTraversal(root)
	fmt.Println(traversal)
}
