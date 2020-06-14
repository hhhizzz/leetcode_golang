package _16

import "sort"

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := int(2E4 + 1)
	resDiff := res - target
	for i := 0; i < len(nums); i++ {
		for l, r := i+1, len(nums)-1; l < r; {
			sum := nums[i] + nums[l] + nums[r]
			diff := sum - target
			if diff == 0 {
				return sum
			} else {
				if abs(diff) < abs(resDiff) {
					resDiff = diff
					res = sum
				}
				if diff < 0 {
					l++
				} else {
					r--
				}
			}
		}
	}
	return res
}
