package _95

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTree(nums []int) []*TreeNode {
	if len(nums) == 0 {
		return nil
	} else if len(nums) == 1 {
		return []*TreeNode{{nums[0], nil, nil}}
	} else {
		var result []*TreeNode
		for i := 0; i < len(nums); i++ {
			left := generateTree(nums[:i])
			right := generateTree(nums[i+1:])
			if len(left) == 0 {
				for _, rightTree := range right {
					result = append(result, &TreeNode{Val: nums[i], Left: nil, Right: rightTree})
				}
			} else if len(right) == 0 {
				for _, leftTree := range left {
					result = append(result, &TreeNode{Val: nums[i], Left: leftTree, Right: nil})
				}
			} else {
				for _, leftTree := range left {
					for _, rightTree := range right {
						result = append(result, &TreeNode{Val: nums[i], Left: leftTree, Right: rightTree})
					}
				}
			}
		}
		return result
	}
}

func generateTrees(n int) []*TreeNode {
	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}
	return generateTree(nums)
}
