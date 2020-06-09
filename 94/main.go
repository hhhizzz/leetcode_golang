package _94

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//比较tricky的写法，更简洁但是稍微用了一些技巧
func inorderTraversal(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode

	current := root

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
		result = append(result, current.Val)

		current = current.Right
	}
	return result
}

//更符合直觉的写法，每次一直往左走到底加入栈中，出栈后变为右孩子然后一直向左
func inorderTraversal1(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode

	current := root
	for current != nil {
		stack = append(stack, current)
		current = current.Left
	}
	for len(stack) != 0 {
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, current.Val)

		current = current.Right
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
	}
	return result
}
