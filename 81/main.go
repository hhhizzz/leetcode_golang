package _81

// 递归写法感觉更易理解一点
func search(nums []int, target int) bool {
	if len(nums) <= 0 {
		return false
	}
	if len(nums) == 1 {
		return target == nums[0]
	}
	mid := len(nums) >> 1
	if nums[mid] == target {
		return true
	}
	if nums[0] < nums[mid] {
		//左边有序
		if target < nums[mid] && target >= nums[0] {
			return search(nums[:mid], target)
		} else {
			return search(nums[mid+1:], target)
		}
	} else if nums[mid] < nums[len(nums)-1] {
		//右边有序
		if target > nums[mid] && target <= nums[len(nums)-1] {
			return search(nums[mid+1:], target)
		} else {
			return search(nums[:mid], target)
		}
	} else {
		//两边都相等的情况，两边都判断
		return search(nums[:mid], target) || search(nums[mid+1:], target)
	}
}

// 非递归写法
func search2(nums []int, target int) bool {
	l := 0
	r := len(nums)
	for l < r {
		m := l + (r-l)>>1
		if nums[m] == target {
			return true
		} else {
			if nums[l] < nums[m] {
				// 左边有序
				if target < nums[m] && target >= nums[l] {
					r = m
				} else {
					l = m + 1
				}
			} else if nums[m] < nums[r-1] {
				// 右边有序
				if target > nums[m] && target <= nums[r-1] {
					l = m + 1
				} else {
					r = m
				}
			} else {
				if nums[r-1] == target {
					return true
				} else {
					r = r - 1
				}
			}
		}
	}
	return false
}
