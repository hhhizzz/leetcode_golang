package _106

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(postorder)
	if n == 0 {
		return nil
	}
	value := postorder[n-1]
	root := &TreeNode{Val: value}
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == value {
			root.Left = buildTree(inorder[:i], postorder[:i])
			root.Right = buildTree(inorder[i+1:], postorder[i:n-1])
			break
		}
	}

	return root
}
