package _663

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func dfs(root *TreeNode, target int, result *bool) int {
    left := 0
    right := 0
    if root.Left != nil {
        left = dfs(root.Left, target, result)
        if left*2 == target {
            *result = true
        }
    }
    if *result {
        return -1
    }
    if root.Right != nil {
        right = dfs(root.Right, target, result)
        if right*2 == target {
            *result = true
        }
    }
    if *result {
        return -1
    }
    return root.Val + left + right
}

func checkEqualTree(root *TreeNode) bool {
    queue := []*TreeNode{root}
    sum := 0
    for len(queue) != 0 {
        current := queue[0]
        queue = queue[1:]

        sum += current.Val
        if current.Left != nil {
            queue = append(queue, current.Left)
        }
        if current.Right != nil {
            queue = append(queue, current.Right)
        }
    }
    result := false

    dfs(root, sum, &result)

    return result
}
