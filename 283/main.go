package _283

//看起来很简单的一道题，其实还是有点考验编码能力的
func moveZeroes(nums []int) {
	pivot := len(nums)
	for i := 0; i < pivot; i++ {
		if nums[i] == 0 {
			pivot--
			for j := i; j < pivot; j++ {
				nums[j] = nums[j+1]
			}
			nums[pivot] = 0
			i--
		}
	}
}
