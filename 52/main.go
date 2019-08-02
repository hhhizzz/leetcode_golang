package _52

func dfs(row, col, pie, na int, cnt *int, n *int) {
    if row >= *n {
        *cnt++
        return
    }
    bits := ^(col | pie | na) & ((1 << uint(*n)) - 1)
    for bits != 0 {
        p := bits & -bits
        dfs(row+1, col|p, (pie|p)<<1, (na|p)>>1, cnt, n)
        bits -= p
    }
}

func totalNQueens(n int) int {
    result := 0
    dfs(0, 0, 0, 0, &result, &n)
    return result
}
