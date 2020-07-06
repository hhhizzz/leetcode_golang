package _53

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func max3(a, b, c int) int {
	return max(max(a, b), c)
}

// 直接贪心
func maxSubArray(nums []int) int {
	// 由于题目说子串至少要一个元素，因此不用判断nums为空的情况
	sum := nums[0]
	res := sum
	for i := 1; i < len(nums); i++ {
		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		res = max(res, sum)
	}
	return res
}

// 四个值分别是 通过左端点的最大和，整体最大和，通过右端点的最大和，以及整体和
func combine(nums []int) (int, int, int, int) {
	if len(nums) == 0 {
		return 0, 0, 0, 0
	} else if len(nums) == 1 {
		return nums[0], nums[0], nums[0], nums[0]
	} else {
		mid := len(nums) >> 1
		left1, left2, left3, leftSum := combine(nums[:mid])
		right1, right2, right3, rightSum := combine(nums[mid:])
		left := max(left1, leftSum+right1)
		center := max3(left2, right2, left3+right1)
		right := max(left3+rightSum, right3)
		sum := leftSum + rightSum
		return left, center, right, sum
	}
}

//按照题目要求的分而治之
func maxSubArray2(nums []int) int {
	_, center, _, _ := combine(nums)
	return center
}
