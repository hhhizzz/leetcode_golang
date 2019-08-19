package _39

func dfs(result *[][]int, nums *[]int, current *[]int, index int, target int) {
    defer func() {
        if len(*current) >= 1 {
            *current = (*current)[:len(*current)-1]
        }
    }()
    if target == 0 {
        r := make([]int, len(*current))
        copy(r, *current)
        *result = append(*result, r)
        return
    }
    if target < 0 {
        return
    }
    for i := index; i < len(*nums); i++ {
        *current = append(*current, (*nums)[i])
        dfs(result, nums, current, i, target-(*nums)[i])
    }
}

func combinationSum(candidates []int, target int) [][]int {
    result := make([][]int, 0)
    current := make([]int, 0)
    dfs(&result, &candidates, &current, 0, target)
    return result
}
