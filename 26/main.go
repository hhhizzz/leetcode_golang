package _26

func removeDuplicates(nums []int) int {
	length := 0
	for i := 0; i < len(nums); i++ {
		nums[length] = nums[i]
		length++
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}
	}
	return length
}
