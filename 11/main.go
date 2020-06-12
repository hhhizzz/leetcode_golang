package _11

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	res := 0
	for i, j := 0, len(height)-1; i != j; {
		h := min(height[i], height[j])
		res = max(res, h*(j-i))
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return res
}
