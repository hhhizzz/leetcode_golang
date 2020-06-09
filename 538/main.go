package _538

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inOrder(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	inOrder(root.Left, sum)
	*sum -= root.Val
	root.Val += *sum
	inOrder(root.Right, sum)
}

func findSum(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	*sum += root.Val
	findSum(root.Left, sum)
	findSum(root.Right, sum)
}

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	findSum(root, &sum)
	inOrder(root, &sum)
	return root
}
