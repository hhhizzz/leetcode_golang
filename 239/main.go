package _239

//双端队列法 window是双端队列，第一个元素表示当前位置到k之前的位置最大的元素，每次往后迭代的时候，新进的元素向队列左端推进直到到前一个元素比它大或者到了最先一个元素
func maxSlidingWindow2(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	output := make([]int, 0)
	window := make([]int, 0)
	for i, x := range nums {
		if len(window) > 0 && window[0] <= i-k {
			window = window[1:]
		}
		for len(window) > 0 && nums[window[len(window)-1]] < x {
			window = window[:len(window)-1]
		}
		window = append(window, i)
		if i >= k-1 {
			output = append(output, nums[window[0]])
		}
	}
	return output
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

//方法1： 构建两个数组，分别表示每隔k的范围内最大值，最后出结果的时候找需要范围的两边即可
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nums
	}
	output := make([]int, 0)
	left, right := make([]int, len(nums)), make([]int, len(nums))
	left[0], right[len(nums)-1] = nums[0], nums[len(nums)-1]

	for i, j := 1, len(nums)-2; i < len(nums); i, j = i+1, j-1 {
		if i%k == 0 {
			left[i] = nums[i]
		} else {
			left[i] = max(nums[i], left[i-1])
		}
		if j%k == 0 {
			right[j] = nums[j]
		} else {
			right[j] = max(nums[j], right[j+1])
		}
	}
	for i := 0; i+k-1 < len(nums); i++ {
		result := max(right[i], left[i+k-1])
		output = append(output, result)
	}
	return output
}
