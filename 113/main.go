package _113

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, sum, target int, path []int, result *[][]int) {
	if root == nil {
		return
	}
	sum += root.Val
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil && sum == target {
		newArray := make([]int, len(path))
		copy(newArray, path)
		*result = append(*result, newArray)
	} else {
		dfs(root.Left, sum, target, path, result)
		dfs(root.Right, sum, target, path, result)
	}
}

func pathSum(root *TreeNode, sum int) [][]int {
	var result [][]int
	dfs(root, 0, sum, []int{}, &result)
	return result
}
