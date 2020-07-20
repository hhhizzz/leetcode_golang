package _84

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 原理是遍历每个矩形，找到这个矩形左右最近比它矮的矩形，例如heights[left] < heights[i] && heights[i] < heights[right]
// 那么这个矩形高度能形成的最大矩形面积就是不包括left right的宽度乘以heights[i]
// 使用单调栈优化，一个单调递增的栈，每次出栈的时候 栈顶元素就是出栈元素的left，当前遍历元素就是right
func largestRectangleArea(heights []int) int {
	var stack []int
	var result int
	// 最后放一个-1以确保最终栈能被完全弹出
	heights = append(heights, -1)
	for i := 0; i < len(heights); i++ {
		if len(stack) == 0 || heights[stack[len(stack)-1]] < heights[i] {
			stack = append(stack, i)
		} else {
			// 这里其实相等也可以不出栈不影响结果，不过让单调栈严格单调更合理
			for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				right := i
				if len(stack) > 0 {
					left := stack[len(stack)-1]
					wide := right - left - 1
					current := wide * heights[top]
					result = max(current, result)
				} else {
					wide := right
					current := wide * heights[top]
					result = max(current, result)
				}
			}
			stack = append(stack, i)
		}
	}
	return result
}
