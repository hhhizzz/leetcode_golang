package _80

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last := nums[0]
	cnt := 1
	pos := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == last {
			cnt += 1
		} else {
			last = nums[i]
			cnt = 1
		}
		if cnt <= 2 {
			nums[pos] = nums[i]
			pos++
		}
	}
	return pos
}
