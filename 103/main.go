package _103

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNodeWithLevel struct {
	*TreeNode
	Level int
}

func reverse(array []int) {
	for i := 0; i < len(array)>>1; i++ {
		array[i], array[len(array)-1-i] = array[len(array)-1-i], array[i]
	}
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	stack := []*TreeNodeWithLevel{
		{root, 1},
	}
	currentLevel := 1
	var currentLevelArray []int
	for len(stack) != 0 {
		current := stack[0]
		stack = stack[1:]
		if current.Left != nil {
			stack = append(stack, &TreeNodeWithLevel{current.Left, current.Level + 1})
		}
		if current.Right != nil {
			stack = append(stack, &TreeNodeWithLevel{current.Right, current.Level + 1})
		}
		if current.Level == currentLevel {
			currentLevelArray = append(currentLevelArray, current.Val)
		} else {
			newArray := make([]int, len(currentLevelArray))
			copy(newArray, currentLevelArray)
			if currentLevel%2 == 0 {
				reverse(newArray)
			}
			result = append(result, newArray)
			currentLevelArray = []int{current.Val}
			currentLevel += 1
		}
	}
	if len(currentLevelArray) != 0 {
		if currentLevel%2 == 0 {
			reverse(currentLevelArray)
		}
		result = append(result, currentLevelArray)
	}
	return result
}
