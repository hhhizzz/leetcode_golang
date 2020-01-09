package _437

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func dfs(root *TreeNode, sum int, result *int, parentSum []int) {
    if root == nil {
        return
    }
    if root.Val == sum {
        *result++
    }
    nextParentSum := make([]int, len(parentSum))
    copy(nextParentSum, parentSum)
    //fmt.Printf("%d: %v\n",root.Val,parentSum)
    for i := 0; i < len(parentSum); i++ {
        if root.Val+nextParentSum[i] == sum {
            *result++
        }
        nextParentSum[i] += root.Val
    }

    nextParentSum = append(nextParentSum, root.Val)

    dfs(root.Left, sum, result, nextParentSum)
    dfs(root.Right, sum, result, nextParentSum)

}

func pathSum(root *TreeNode, sum int) int {
    result := 0
    dfs(root, sum, &result, []int{})
    return result
}
