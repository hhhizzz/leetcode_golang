package _152

func maxNum(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		} else {
			return c
		}
	} else {
		if b > c {
			return b
		} else {
			return c
		}
	}
}
func minNum(a, b, c int) int {
	if a < b {
		if c < a {
			return c
		} else {
			return a
		}
	} else {
		if c < b {
			return c
		} else {
			return b
		}
	}
}

func maxProduct(nums []int) int {
	max := make([]int, len(nums))
	min := make([]int, len(nums))
	max[0], min[0] = nums[0], nums[0]
	res := nums[0]

	for i := 1; i < len(nums); i++ {
		max[i] = maxNum(max[i-1]*nums[i], min[i-1]*nums[i], nums[i])
		min[i] = minNum(max[i-1]*nums[i], min[i-1]*nums[i], nums[i])
		if max[i] > res {
			res = max[i]
		}
	}

	return res
}
