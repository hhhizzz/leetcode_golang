package _35

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
