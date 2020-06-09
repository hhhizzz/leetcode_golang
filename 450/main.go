package _450

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//找到某节点的中序后继并删除这个节点
func succ(root *TreeNode) int {
	current := root.Right
	parent := root
	for current.Left != nil {
		parent = current
		current = current.Left
	}
	//中序后继必然无左子树，因此删除只需用右子树代替自己
	if parent != root {
		parent.Left = current.Right
	} else {
		parent.Right = current.Right
	}
	return current.Val
}

//递归删除方案
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		//如果命中，检查是否有某个子树是空的
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		//如果没有某个子树是空的，查找中序后继并替换
		root.Val = succ(root)
	} else {
		//向下查找
		if key < root.Val {
			root.Left = deleteNode(root.Left, key)
		} else {
			root.Right = deleteNode(root.Right, key)
		}
	}
	return root
}

func deleteChild(child *TreeNode) *TreeNode {
	if child.Left == nil {
		return child.Right
	}
	if child.Right == nil {
		return child.Left
	}
	child.Val = succ(child)
	return child
}

//非递归的方案，内存用量没有减少，速度还慢了。。。
func deleteNode1(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == key {
		return deleteChild(root)
	} else {
		current := root
		for current != nil {
			if current.Left != nil && current.Left.Val == key {
				current.Left = deleteChild(current.Left)
				break
			}
			if current.Right != nil && current.Right.Val == key {
				current.Right = deleteChild(current.Right)
				break
			}
			if key < current.Val {
				current = current.Left
			} else if key > current.Val {
				current = current.Right
			}
		}
		return root
	}
}
