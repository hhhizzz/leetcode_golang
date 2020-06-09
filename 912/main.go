package _912

//这道题给定了数据的大小范围，明显是想考察基数排序
func sortArray(nums []int) []int {
	temp := make([]int, 100001)
	for i := 0; i < len(nums); i++ {
		temp[nums[i]+50000] += 1
	}
	result := make([]int, len(nums))
	idx := 0
	for i := 0; i < len(temp); i++ {
		for j := 0; j < temp[i]; j++ {
			result[idx] = i - 50000
			idx++
		}
	}
	return result
}
