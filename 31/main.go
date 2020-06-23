package _31

import "sort"

func nextPermutation(nums []int) {
	// 1 2 3 4 5
	// 1 2 3 5 4
	// 1 2 4 3 5
	// 1 2 4 5 3
	// 1 2 5 3 4
	// 1 2 5 4 3
	// 1 3 2 4 5
	// 1 3 2 5 4
	// ...
	// 观察规律，如果一个后缀是倒序的，说明这个后缀已经完全遍历完了，不会有下一个序列了，如果这个后缀前面有数，那么可以从后缀里找出最小的一个比前面这个数大的交换
	// 然后把这个后缀重排序即可
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			for j := i + 1; j < len(nums); j++ {
				if j == len(nums)-1 || nums[j+1] <= nums[i] {
					nums[i], nums[j] = nums[j], nums[i]
					sort.Ints(nums[i+1:])
					return
				}
			}
		}
	}
	for i := 0; i < len(nums)>>1; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}
