package _46

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
            dfs(j, current)
        }
    }
}

func permute(nums []int) [][]int {
    array = nums
    used = make([]bool, len(nums))
    result = make([][]int, 0)
    for i := 0; i < len(nums); i++ {
        current := make([]int, 0)
        dfs(i, &current)
    }
    return result
}
