package _15

func qsort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	pivot := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[0] {
			pivot++
			nums[i], nums[pivot] = nums[pivot], nums[i]
		}
	}
	nums[pivot], nums[0] = nums[0], nums[pivot]
	qsort(nums[:pivot])
	qsort(nums[pivot+1:])
}

//方案1，使用map
func threeSum1(nums []int) [][]int {
	var result [][]int

	qsort(nums)

	res := map[[3]int]bool{}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		m := map[int]bool{}
		for j := i + 1; j < len(nums); j++ {
			if _, ok := m[nums[j]]; ok {
				array := [3]int{nums[i], -nums[i] - nums[j], nums[j]}
				if _, ok := res[array]; !ok {
					result = append(result, array[:])
					res[array] = true
				}
			} else {
				m[-nums[i]-nums[j]] = true
			}
		}

	}
	return result
}

func threeSum(nums []int) [][]int {
	qsort(nums)
	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			total := nums[i] + nums[left] + nums[right]
			if total > 0 {
				right--
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if total < 0 {
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				right--
				for left < right && nums[right] == nums[right+1] {
					right--
				}
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			}
		}
	}
	return result
}
