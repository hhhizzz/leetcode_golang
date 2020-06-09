package _209

import "math"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minSubArrayLen(s int, nums []int) int {
	left := 0
	right := 0
	sum := 0
	result := math.MaxInt32
	for right < len(nums) {
		sum += nums[right]
		if sum >= s {
			result = min(result, right-left+1)
			for left < right {
				sum -= nums[left]
				left++
				if sum >= s {
					result = min(result, right-left+1)
				} else {
					break
				}
			}
		}
		right++
	}
	if result == math.MaxInt32 {
		return 0
	}
	return result
}
