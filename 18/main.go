package _18

func qsort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	pivot := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[0] {
			pivot++
			nums[pivot], nums[i] = nums[i], nums[pivot]
		}
	}
	nums[0], nums[pivot] = nums[pivot], nums[0]

	qsort(nums[:pivot])
	qsort(nums[pivot+1:])
}

//先尝试用map的方法，因为要大量申请数组，用go的话特别慢，但其实和双指针法都是O(n^3)的
func fourSum(nums []int, target int) [][]int {
	qsort(nums)
	var result [][]int
	res := map[[4]int]bool{}

	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			m := map[int]bool{}
			for k := j + 1; k < len(nums); k++ {
				goal := target - nums[i] - nums[j] - nums[k]
				if _, ok := m[nums[k]]; ok {
					array := [4]int{nums[i], nums[j], goal, nums[k]}
					if _, ok := res[array]; !ok {
						result = append(result, array[:])
						res[array] = true
					}
				} else {
					m[goal] = true
				}
			}
		}
	}
	return result
}
