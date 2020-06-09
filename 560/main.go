package _560

//利用sum数组暴力找
func subarraySum1(nums []int, k int) int {
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}
	result := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			var sum int
			if i == 0 {
				sum = sums[j]
			} else {
				sum = sums[j] - sums[i-1]
			}
			if sum == k {
				result++
			}
		}
	}
	return result
}

//使用map存当前从0加到某个位置的和出现次数
func subarraySum(nums []int, k int) int {
	m := map[int]int{}
	m[0] = 1
	count := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if v, ok := m[sum-k]; ok {
			count += v
		}
		m[sum] += 1
	}
	return count
}
