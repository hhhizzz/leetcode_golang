package _33

func search(nums []int, target int) int {
	left := 0
	right := len(nums)
	// 稍微需要思考一下的二分，重要一点就是至少有一个区间是单调的，找出这个区间判断target是否在这个区间即可
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		} else {
			// 如果最左边的数小于中间数，那么左边一定是单调的，下同
			if nums[left] < nums[mid] {
				if target >= nums[left] && target < nums[mid] {
					right = mid
				} else {
					left = mid + 1
				}
			} else {
				if target > nums[mid] && target <= nums[right-1] {
					left = mid + 1
				} else {
					right = mid
				}
			}
		}
	}
	return -1
}
