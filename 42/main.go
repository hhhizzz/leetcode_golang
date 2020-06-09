package _42

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//此题主要要理解的地方在于计算一个点能够积攒的水，方法是取左右最高度的最小值减去本地的高度
func trap(height []int) int {
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(leftMax[i-1], height[i-1])
	}
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i+1])
	}
	result := 0
	for i := 0; i < len(height); i++ {
		side := min(leftMax[i], rightMax[i])
		if side <= height[i] {
			continue
		}
		current := side - height[i]
		result += current
	}
	return result
}
