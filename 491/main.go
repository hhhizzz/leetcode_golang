package _491

func dfs(nums []int, i int, result *[][]int, current []int) {
	showed := map[int]bool{}
	for j := i + 1; j < len(nums); j++ {
		if nums[j] >= nums[i] {
			if showed[nums[j]] {
				continue
			}
			showed[nums[j]] = true
			current = append(current, nums[j])
			newArray := make([]int, len(current))
			copy(newArray, current)
			*result = append(*result, newArray)
			dfs(nums, j, result, current)
			current = current[:len(current)-1]
		}
	}
}

func findSubsequences(nums []int) [][]int {
	var result [][]int
	showed := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if showed[nums[i]] {
			continue
		}
		showed[nums[i]] = true
		dfs(nums, i, &result, []int{nums[i]})
	}
	return result
}
