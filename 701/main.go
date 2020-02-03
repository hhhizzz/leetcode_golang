package _701

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{val, nil, nil}
    }
    current := root
    var parent *TreeNode
    for current != nil {
        parent = current
        if val < current.Val {
            current = current.Left
        } else if current.Val < val {
            current = current.Right
        }
    }
    if val < parent.Val {
        parent.Left = &TreeNode{val, nil, nil}
    } else {
        parent.Right = &TreeNode{val, nil, nil}
    }
    return root
}
