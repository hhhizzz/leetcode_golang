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

func kthSmallest1(root *TreeNode, k int) int {
    var result int
    dfs(root, &k, &result)
    return result
}

//非递归中序的方案
func kthSmallest(root *TreeNode, k int) int {
    var stack []*TreeNode
    current := root
    for current != nil {
        stack = append(stack, current)
        current = current.Left
    }
    for len(stack) != 0 {
        current := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        k--

        //fmt.Println(current.Val)
        if k == 0 {
            return current.Val
        }

        current = current.Right
        for current != nil {
            stack = append(stack, current)
            current = current.Left
        }
    }
    return -1
}
