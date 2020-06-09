package _257

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(realRoot bool, root *TreeNode, current string, result *[]string) {
	if root == nil {
		return
	}
	if !realRoot {
		current += "->"
	}

	current += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		*result = append(*result, current)
	}
	dfs(false, root.Left, current, result)
	dfs(false, root.Right, current, result)
}

func binaryTreePaths(root *TreeNode) []string {
	var result []string
	dfs(true, root, "", &result)
	return result
}
