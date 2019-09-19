package _102

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
    queue := []*TreeNode{root}
    levelQueue := []int{0}
    var result [][]int
    if root == nil {
        return result
    }
    for len(queue) != 0 {
        current := queue[0]
        currentLevel := levelQueue[0]
        if len(result) <= currentLevel {
            result = append(result, []int{current.Val})
        } else {
            result[currentLevel] = append(result[currentLevel], current.Val)
        }
        queue = queue[1:]
        levelQueue = levelQueue[1:]
        if current.Left != nil {
            queue = append(queue, current.Left)
            levelQueue = append(levelQueue, currentLevel+1)
        }
        if current.Right != nil {
            queue = append(queue, current.Right)
            levelQueue = append(levelQueue, currentLevel+1)
        }
    }
    return result
}
