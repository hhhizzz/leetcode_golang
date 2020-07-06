package _55

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 方法1 倒着走看能不能走回去
func canJump(nums []int) bool {
	lastPos := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i]+i >= lastPos {
			lastPos = i
		}
	}
	return lastPos == 0
}

// 方法2 正着走尝试最远能走多远
func canJump2(nums []int) bool {
	end := 0
	for i := 0; i < len(nums); i++ {
		if i > end {
			return false
		}
		end = max(end, i+nums[i])
	}
	return end >= len(nums)-1
}
