package _333

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func min2(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max2(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min3(a, b, c int) int {
    return min2(min2(a, b), c)
}

func max3(a, b, c int) int {
    return max2(max2(a, b), c)
}

func dfs(root *TreeNode) (max, min, num, result int) {
    max = math.MinInt32
    min = math.MaxInt32
    num = 0
    result = 0
    if root == nil {
        return
    }
    leftMax, leftMin, leftNum, leftResult := dfs(root.Left)
    rightMax, rightMin, rightNum, rightResult := dfs(root.Right)

    max = max3(leftMax, root.Val, rightMax)
    min = min3(leftMin, root.Val, rightMin)
    result = max2(leftResult, rightResult)

    if leftNum == leftResult && rightNum == rightResult && leftMax < root.Val && root.Val < rightMin {
        num = leftNum + rightNum + 1
        result = num
        return
    }
    return
}

func largestBSTSubtree(root *TreeNode) int {
    _, _, _, result := dfs(root)
    return result
}
