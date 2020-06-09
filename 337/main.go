package _337

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(root *TreeNode) int {
	rob, notRob := helper(root)
	return max(rob, notRob)
}

//第一个返回值表示抢root，第二个表示不抢
func helper(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	leftRob, leftNotRob := helper(root.Left)
	rightRob, rightNotRob := helper(root.Right)

	return root.Val + leftNotRob + rightNotRob, max(leftRob, leftNotRob) + max(rightRob, rightNotRob)
}
