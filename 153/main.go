package _153

//肯定不能全部扫，目标是要达到O(logN)
//方法是找“山谷”，就是两边值都比自己大的值，可以认为nums[len(nums)]=+∞
func findMin(nums []int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := left + ((right - left) >> 1)
		if (mid == 0 || nums[mid] < nums[mid-1]) && (mid == len(nums)-1 || nums[mid] < nums[mid+1]) {
			//如果找到直接返回
			return nums[left]
		} else if nums[mid] <= nums[right-1] {
			//如果right值大于等于mid，说明右边值单调的，往左边找（=的情况仅出现在mid==right-1）
			right = mid
		} else {
			//如果右边没有单调那么结果一定在右边
			left = mid + 1
		}
	}
	return nums[left]
}
