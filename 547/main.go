package _547

func dfs(M *[][]int, found *[]bool, i int) {
	if (*found)[i] {
		return
	}
	(*found)[i] = true
	for j := 0; j < len((*M)[i]); j++ {
		if i != j && (*M)[i][j] == 1 {
			dfs(M, found, j)
		}
	}
}

func findCircleNum(M [][]int) int {
	found := make([]bool, len(M))
	count := 0
	for i := 0; i < len(M); i++ {
		if !found[i] {
			count++
			dfs(&M, &found, i)
		}
	}
	return count
}
