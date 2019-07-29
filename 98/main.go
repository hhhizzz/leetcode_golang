package _98

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func isValid(root *TreeNode, min, max *int) bool {
    if root == nil {
        return true
    }
    if min != nil && root.Val <= *min {
        return false
    }
    if max != nil && root.Val >= *max {
        return false
    }
    return isValid(root.Left, nil, &root.Val) && isValid(root.Right, &root.Val, nil)
}

func isValidBST(root *TreeNode) bool {
    return isValid(root, nil, nil)
}
