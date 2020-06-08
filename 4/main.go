package _4

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 基本方针是不断从中减少目标长度的一半直到找到结果
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	// 根据长度的奇偶性判断应该找到哪个位置的值
	if length%2 == 1 {
		return float64(findKth(nums1, nums2, length>>1+1))
	} else {
		left := findKth(nums1, nums2, length>>1)
		right := findKth(nums1, nums2, length>>1+1)
		return float64(left+right) / 2.0
	}
}

// 找到 nums1 和 nums2 合起来排序后排第k的值，k从1开始数
func findKth(nums1, nums2 []int, k int) int {
	// 首先确认nums1是比较短的
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	// 递归尾部，当k只剩1的时候，可以返回结果了
	if k == 1 {
		if len(nums1) == 0 {
			return nums2[0]
		} else {
			return min(nums1[0], nums2[0])
		}
	} else {
		// 否则在两个数组中各自查找 k / 2位置的值
		if len(nums1) == 0 {
			// 边界情况，如果nums1为空，直接返回nums2的结果
			return nums2[k-1]
		}
		// nums1比较短，有可能越界，注意这里是len(nums1)因为这里k k21 k22都是从1开始数的，也就是表示开区间
		k21 := min(len(nums1), k>>1)
		k22 := k >> 1
		//原理见上，检查数字的时候需要-1
		if nums1[k21-1] < nums2[k22-1] {
			// 如果第一个数组的k / 2 位置的值比较小，说明 nums[:k21]这段的所有数一定是比目标值小的，可以排除掉，然后用递归查找剩下的数组
			return findKth(nums1[k21:], nums2, k-k21)
		} else {
			return findKth(nums1, nums2[k22:], k-k22)
		}
	}
}
