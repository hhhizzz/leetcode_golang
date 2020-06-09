package _1150

//其实很简单的题，稍微练一下二分
func isMajorityElement(nums []int, target int) bool {
	left := 0
	right := len(nums)

	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			for left = mid; left > 0 && nums[left] == target; left-- {

			}
			if nums[left] != target {
				left++
			}
			for right = mid; right < len(nums) && nums[right] == target; right++ {

			}
			break
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	//fmt.Println(left, right)
	//记得二分失败的情况是left==len(nums)
	if left >= len(nums) {
		return false
	}
	return right-left > len(nums)/2
}
