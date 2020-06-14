package _15

import "sort"

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
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-2; i++ {
		for i > 0 && i < len(nums)-1 && nums[i] == nums[i-1] {
			i++
		}
		for l, r := i+1, len(nums)-1; l < r; {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				r--
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}
