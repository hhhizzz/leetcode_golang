package _448

func findDisappearedNumbers(nums []int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		for nums[nums[i]-1] != nums[i] && nums[i] != i+1 {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	//fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			result = append(result, i+1)
		}
	}
	return result
}
