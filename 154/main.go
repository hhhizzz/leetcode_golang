package _154

func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1
	for l < r {
		m := l + (r-l)>>1
		if nums[m] < nums[r] {
			r = m
		} else if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = r - 1
		}
	}
	return nums[l]
}
