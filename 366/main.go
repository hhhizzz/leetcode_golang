package _366

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func dfs(root *TreeNode, leaves *[]int) bool {
    if root.Left == nil && root.Right == nil {
        *leaves = append(*leaves, root.Val)
        return true
    }
    if root.Left != nil {
        if dfs(root.Left, leaves) {
            root.Left = nil
        }
    }
    if root.Right != nil {
        if dfs(root.Right, leaves) {
            root.Right = nil
        }
    }
    return false
}

func findLeaves(root *TreeNode) [][]int {
    var result [][]int
    if root == nil {
        return result
    }
    for {
        var leaves []int
        isEmpty := dfs(root, &leaves)
        result = append(result, leaves)
        if isEmpty {
            return result
        }
    }
}
