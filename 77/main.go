package _77

func dfs(array *[]int, current *[]int, result *[][]int, k int) {
    if k == 0 {
        newResult := make([]int, len(*current))
        copy(newResult, *current)
        *result = append(*result, newResult)
        return
    } else {
        for i := 0; i < len(*array); i++ {
            *current = append(*current, (*array)[i])
            newArray := (*array)[i+1:]
            dfs(&newArray, current, result, k-1)
            *current = (*current)[:len(*current)-1]
        }
    }
}

func combine(n int, k int) [][]int {
    array := make([]int, n)
    for i := 1; i <= n; i++ {
        array[i-1] = i
    }
    var current []int
    var result [][]int
    dfs(&array, &current, &result, k)
    return result
}
