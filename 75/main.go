package _75

func sortColors(nums []int) {
	left := 0
	right := len(nums) - 1
	for right >= left && nums[right] == 2 {
		right--
	}
	for i := 0; i <= right; i++ {
		if nums[i] == 0 {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		} else if nums[i] == 2 {
			nums[right], nums[i] = nums[i], nums[right]
			for right > left && nums[right] == 2 {
				right--
			}
			i--
		}
	}
}
