package _80

func deleteArray(nums *[]int, start, end int) {
	length := end - start
	for i := start; i+length < len(*nums); i++ {
		(*nums)[i] = (*nums)[i+length]
	}
	*nums = (*nums)[:len(*nums)-length]
}

func removeDuplicates(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if i+1 < len(nums) && nums[i] == nums[i+1] {
			j := i + 2
			for ; j < len(nums) && nums[j] == nums[i]; j++ {
			}
			if j > i+2 {
				deleteArray(&nums, i+2, j)
			}
		}
	}
	return len(nums)
}
