package _238

func productExceptSelf(nums []int) []int {
	left := 1
	right := 1
	result := make([]int, len(nums))
	for i := 0; i < len(result); i++ {
		result[i] = 1
	}

	for i := 0; i < len(result); i++ {
		result[i] *= left
		left *= nums[i]
		result[len(result)-1-i] *= right
		right *= nums[len(result)-1-i]
	}

	return result
}
