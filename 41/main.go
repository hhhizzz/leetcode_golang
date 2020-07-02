package _41

// 比较难想到，但是算法不复杂
// 基本思路是把符合条件的数移动到自己的位置上去，然后第二次扫描检查找出第一个不符合自己位置的数
// 本题由于是找正整数，因此需要做一下位置映射，
func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		nums[i]--
	}
	for i := 0; i < len(nums); i++ {
		for nums[i] >= 0 && nums[i] < len(nums) && nums[i] != i && nums[i] != nums[nums[i]] {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i + 1
		}
	}
	return len(nums) + 1
}
