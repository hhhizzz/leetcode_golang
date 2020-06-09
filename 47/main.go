package _47

var array []int
var used []bool
var result [][]int

func dfs(i int, current *[]int) {
	if used[i] {
		return
	}
	used[i] = true
	*current = append(*current, array[i])

	defer func() {
		used[i] = false
		*current = (*current)[:len(*current)-1]
	}()

	if len(*current) == len(array) {
		new := make([]int, len(array))
		copy(new, *current)
		result = append(result, new)
	} else {
		for j := 0; j < len(array); j++ {
			if j > 0 && !used[j-1] && array[j] == array[j-1] {
				continue
			}
			dfs(j, current)
		}
	}
}

func qsort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	pivot := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[0] {
			pivot++
			nums[pivot], nums[i] = nums[i], nums[pivot]
		}
	}
	nums[pivot], nums[0] = nums[0], nums[pivot]
	qsort(nums[:pivot])
	qsort(nums[pivot+1:])
}

func permuteUnique(nums []int) [][]int {
	array = nums
	qsort(array)
	used = make([]bool, len(nums))
	result = make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i != 0 && array[i] == array[i-1] {
			continue
		}
		current := make([]int, 0)
		dfs(i, &current)
	}
	return result
}
