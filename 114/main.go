package _114

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

//后序遍历，找到左子树的最右端点，连接到右子树上
func helper(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    left := helper(root.Left)
    right := helper(root.Right)

    if left != nil {
        left.Right = root.Right
        root.Right = root.Left
        root.Left = nil
    }

    if right != nil {
        return right
    }
    if left != nil {
        return left
    }
    return root
}

//递归法
func flatten(root *TreeNode) {
    helper(root)
}
