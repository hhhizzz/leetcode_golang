package _1

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		complement := target - nums[i]
		if value, ok := m[complement]; ok {
			if i < value {
				return []int{i, value}
			} else {
				return []int{value, i}
			}
		}
		m[num] = i
	}
	return []int{}
}
