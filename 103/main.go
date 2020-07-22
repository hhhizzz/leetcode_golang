package _103

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 记录每层的节点，然后按照不同的顺序写入队列
func zigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	var queue []*TreeNode
	var level []*TreeNode
	level = append(level, root)
	currentLevel := -1

	for len(queue) != 0 || len(level) != 0 {
		if len(queue) == 0 {
			currentLevel++
			// 因为每一层和上一层的数据是反向的，因此是反向填入队列
			for i := len(level) - 1; i >= 0; i-- {
				queue = append(queue, level[i])
			}
			level = []*TreeNode{}
			result = append(result, []int{})
		}
		current := queue[0]
		queue = queue[1:]
		result[currentLevel] = append(result[currentLevel], current.Val)

		// 注意每层的顺序，如果本层是从右往左的情况，应该先放入右子树
		if currentLevel%2 == 0 {
			if current.Left != nil {
				level = append(level, current.Left)
			}
			if current.Right != nil {
				level = append(level, current.Right)
			}
		} else {
			if current.Right != nil {
				level = append(level, current.Right)
			}
			if current.Left != nil {
				level = append(level, current.Left)
			}
		}
	}
	return result
}
