package _95

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(l, r int) []*TreeNode {
	if l > r {
		return []*TreeNode{nil}
	}
	if l == r {
		return []*TreeNode{{Val: l}}
	} else {
		var res []*TreeNode
		for i := l; i <= r; i++ {
			left := helper(l, i-1)
			right := helper(i+1, r)
			for _, leftTree := range left {
				for _, rightTree := range right {
					res = append(res, &TreeNode{Val: i, Left: leftTree, Right: rightTree})
				}
			}
		}
		return res
	}
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return helper(1, n)
}
