package _285

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	s := []*TreeNode{}
	current := root
	for current != nil {
		s = append(s, current)
		current = current.Left
	}
	flag := false
	for len(s) != 0 {
		current = s[len(s)-1]
		s = s[:len(s)-1]

		if flag {
			return current
		}
		if !flag && current == p {
			flag = true
		}
		if current.Right != nil {
			current = current.Right
			for current != nil {
				s = append(s, current)
				current = current.Left
			}
		}
	}
	return nil
}
