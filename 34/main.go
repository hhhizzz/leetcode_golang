package _34

func searchRange(nums []int, target int) []int {
	left := 0
	right := len(nums)
	res := []int{-1, -1}
	for left < right {
		mid := (left + right) >> 1
		if target == nums[mid] {
			if mid == 0 || nums[mid-1] != nums[mid] {
				res[0] = mid
				break
			} else {
				right = mid
			}
		} else if target < nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	left = 0
	right = len(nums)
	for left < right {
		mid := (left + right) >> 1
		if target == nums[mid] {
			if mid == len(nums)-1 || nums[mid+1] != nums[mid] {
				res[1] = mid
				break
			} else {
				left = mid + 1
			}
		} else if target < nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return res
}
