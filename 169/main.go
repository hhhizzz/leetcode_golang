package _169

func majorityElement1(nums []int) int {
	m := map[int]int{}
	half := len(nums) >> 1

	for _, num := range nums {
		if _, ok := m[num]; ok {
			m[num] += 1
		} else {
			m[num] = 1
		}
		if m[num] > half {
			return num
		}
	}
	return -1
}

func majorityElement(nums []int) int {
	count := 0
	current := nums[0]
	for i := 0; i < len(nums); i++ {
		if current == nums[i] {
			count++
		} else {
			count--
		}
		if count == 0 {
			current = nums[i+1]
		}
	}
	return current
}
