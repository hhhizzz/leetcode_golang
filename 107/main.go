package _107

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

type NodeWithLevel struct {
    *TreeNode
    Level int
}

func levelOrderBottom(root *TreeNode) [][]int {
    queue := []*NodeWithLevel{{root, 0}}
    var result [][]int
    if root == nil {
        return result
    }
    for len(queue) != 0 {
        current := queue[0]
        if len(result) <= current.Level {
            result = append([][]int{{current.Val}}, result...)
        } else {
            result[0] = append(result[0], current.Val)
        }
        if current.Left != nil {
            queue = append(queue, &NodeWithLevel{current.Left, current.Level + 1})
        }
        if current.Right != nil {
            queue = append(queue, &NodeWithLevel{current.Right, current.Level + 1})
        }
        queue = queue[1:]
    }
    return result
}
