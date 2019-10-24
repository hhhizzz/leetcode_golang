package _230

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func dfs(root *TreeNode, k *int, ans *int) {
    if root == nil {
        return
    }
    dfs(root.Left, k, ans)
    if *k > 1 {
        *k -= 1
    } else if *k == 1 {
        *k -= 1
        *ans = root.Val
        return
    } else {
        return
    }
    dfs(root.Right, k, ans)
}

func kthSmallest(root *TreeNode, k int) int {
    var result int
    dfs(root, &k, &result)
    return result
}
