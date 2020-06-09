package _41

func firstMissingPositive(nums []int) int {
	hasOne := false
	for _, num := range nums {
		if num == 1 {
			hasOne = true
		}
	}
	if !hasOne {
		return 1
	}
	if len(nums) == 1 {
		return 2
	}
	n := len(nums)
	for i, num := range nums {
		if num <= 0 || num > n {
			nums[i] = 1
		}
	}
	bitmap := make([]int, n+1)
	for _, num := range nums {
		bitmap[num] = 1
	}
	for i := 1; i <= n; i++ {
		if bitmap[i] == 0 {
			return i
		}
	}
	return n + 1
}
