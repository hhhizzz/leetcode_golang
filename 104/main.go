package _104

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

type TreeNodeWithLevel struct {
    *TreeNode
    Level int
}

//非递归还是有点麻烦
func maxDepth(root *TreeNode) int {
    var stack []*TreeNodeWithLevel
    result := 0
    current := root
    currentLevel := 1
    for true {
        for current != nil {
            if current.Right != nil {
                stack = append(stack, &TreeNodeWithLevel{current.Right, currentLevel + 1})
            }
            current = current.Left
            currentLevel += 1
        }
        currentLevel -= 1
        if currentLevel > result {
            result = currentLevel
        }
        if len(stack) == 0 {
            break
        }
        current, currentLevel = stack[len(stack)-1].TreeNode, stack[len(stack)-1].Level
        stack = stack[:len(stack)-1]
    }
    return result
}
