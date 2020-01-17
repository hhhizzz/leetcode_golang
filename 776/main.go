package _776

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func splitBST(root *TreeNode, V int) []*TreeNode {
    if root == nil {
        return []*TreeNode{nil, nil}
    } else if root.Val <= V {
        result := splitBST(root.Right, V)
        root.Right = result[0]
        result[0] = root
        return result
    } else {
        result := splitBST(root.Left, V)
        root.Left = result[1]
        result[1] = root
        return result
    }
}
