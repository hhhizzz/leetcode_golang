package _46

func helper(nums []int, used map[int]bool, current []int, res *[][]int) {
	if len(nums) == len(used) {
		temp := make([]int, len(current))
		copy(temp, current)
		*res = append(*res, temp)
	} else {
		for i := 0; i < len(nums); i++ {
			if _, ok := used[i]; ok {
				continue
			}
			current = append(current, nums[i])
			used[i] = true
			helper(nums, used, current, res)
			current = current[:len(current)-1]
			delete(used, i)
		}
	}
}

func permute(nums []int) [][]int {
	var res [][]int
	helper(nums, map[int]bool{}, []int{}, &res)
	return res
}
