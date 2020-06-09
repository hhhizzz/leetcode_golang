package _298

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, parent, length int, result *int) {
	if root == nil {
		return
	}
	if root.Val == parent+1 {
		length += 1
	} else {
		length = 1
	}
	if length > *result {
		*result = length
	}
	dfs(root.Left, root.Val, length, result)
	dfs(root.Right, root.Val, length, result)
}

func longestConsecutive(root *TreeNode) int {
	result := 0
	dfs(root, 0, 0, &result)
	return result
}
