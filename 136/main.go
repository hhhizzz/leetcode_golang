package _136

func singleNumber(nums []int) int {
	a := nums[0]
	for i := 1; i < len(nums); i++ {
		a ^= nums[i]
	}
	return a
}
