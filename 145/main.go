package _145

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//后序的非递归确实很麻烦
func postorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{root}
	var result []int
	if root == nil {
		return result
	}

	current := root
	for len(stack) != 0 {
		//首先判断下一个要取的节点不是当前节点父节点，由于放入栈的方式，每个节点的下一个节点要么是它的右兄弟，要么是父亲节点
		if stack[len(stack)-1].Left != current && stack[len(stack)-1].Right != current {
			//取出下一个节点，按照次序找到它左侧最深节点放入栈中
			current = stack[len(stack)-1]
			for current.Left != nil || current.Right != nil {
				//如果它有左子树，就向左，没有再向右
				if current.Left != nil {
					//先把右孩子放入栈中，因为栈是先入后出
					if current.Right != nil {
						stack = append(stack, current.Right)
					}
					stack = append(stack, current.Left)
				} else if current.Right != nil {
					stack = append(stack, current.Right)
				}
				current = stack[len(stack)-1]
			}
		} //如果下一个要取的节点是当前节点父节点，说明即将要取的节点子树已经被遍历完了，可以进入遍历流程了
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, current.Val)
	}
	return result
}
