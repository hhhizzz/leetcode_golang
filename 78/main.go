package _78

func dfs(nums *[]int, length int, current *[]int, result *[][]int) {
	if length == 0 {
		newArray := make([]int, len(*current))
		copy(newArray, *current)
		*result = append(*result, newArray)
	} else {
		for i := 0; i < len(*nums); i++ {
			*current = append(*current, (*nums)[i])
			nextNum := (*nums)[i+1:]
			dfs(&nextNum, length-1, current, result)
			*current = (*current)[:len(*current)-1]
		}
	}
}

func subsets(nums []int) [][]int {
	var result [][]int
	for i := 0; i <= len(nums); i++ {
		var current []int
		dfs(&nums, i, &current, &result)
	}

	return result
}
