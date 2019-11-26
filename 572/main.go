package _572

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func equal(a, b *TreeNode) bool {
    if a == nil && b == nil {
        return true
    }
    if a == nil && b != nil {
        return false
    }
    if a != nil && b == nil {
        return false
    }
    if a.Val == b.Val {
        return equal(a.Left, b.Left) && equal(a.Right, b.Right)
    } else {
        return false
    }
}
func isSubtree(s *TreeNode, t *TreeNode) bool {
    if s == nil && t != nil {
        return false
    }
    if equal(s, t) {
        return true
    }
    return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}
