package _545

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func boundaryOfBinaryTree(root *TreeNode) []int {
	var stack []*TreeNode
	var result []int

	bound := map[*TreeNode]bool{}

	current := root
	for current != nil {
		if _, ok := bound[current]; !ok {
			result = append(result, current.Val)
			bound[current] = true
		}
		if current != root && current.Left == nil {
			current = current.Right
		} else {
			current = current.Left
		}
	}

	current = root
	for true {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		if len(stack) == 0 {
			break
		}
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if _, ok := bound[current]; !ok && current.Left == nil && current.Right == nil {
			bound[current] = true
			result = append(result, current.Val)
		}
		current = current.Right
	}

	rightBound := []int{}
	current = root
	for current != nil {
		if _, ok := bound[current]; !ok {
			rightBound = append(rightBound, current.Val)
			bound[current] = true
		}
		if current != root && current.Right == nil {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	for i := 0; i < len(rightBound); i++ {
		result = append(result, rightBound[len(rightBound)-1-i])
	}
	return result
}
