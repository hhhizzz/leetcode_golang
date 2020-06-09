package _95

import (
	"fmt"
	"testing"
)

func (node *TreeNode) toString() string {
	if node == nil {
		return ""
	}
	result := fmt.Sprintf("{%d}", node.Val)
	result += node.Left.toString()
	result += node.Right.toString()
	return result
}

func TestGenerateTrees(t *testing.T) {
	var nodes []*TreeNode
	nodes = generateTrees(2)
	for _, node := range nodes {
		fmt.Println(node.toString() + ", ")
	}
	fmt.Println()
	nodes = generateTrees(3)
	for _, node := range nodes {
		fmt.Println(node.toString() + ", ")
	}
	fmt.Println()
	nodes = generateTrees(4)
	for _, node := range nodes {
		fmt.Println(node.toString() + ", ")
	}
	fmt.Println()
}
