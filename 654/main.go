package _654

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func findMax(nums []int) int {
    max := 0
    maxNum := nums[max]
    for i := 1; i < len(nums); i++ {
        if nums[i] > maxNum {
            max = i
            maxNum = nums[max]
        }
    }
    return max
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
    if len(nums) <= 0 {
        return nil
    }
    root := &TreeNode{}
    max := findMax(nums)
    root.Val = nums[max]
    root.Left = constructMaximumBinaryTree(nums[:max])
    root.Right = constructMaximumBinaryTree(nums[max+1:])
    return root
}
