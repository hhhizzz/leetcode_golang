package _100

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil || q == nil {
        if p == nil && q == nil {
            return true
        } else {
            return false
        }
    }

    if p.Val != q.Val {
        return false
    }
    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
