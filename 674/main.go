package _674

func findLengthOfLCIS(nums []int) int {
	result := 0
	current := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] > nums[i-1] {
			current += 1
			if current > result {
				result = current
			}
		} else {
			current = 1
		}
	}
	return result
}
