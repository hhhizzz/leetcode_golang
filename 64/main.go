package _64

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func minPathSum(grid [][]int) int {
    x := len(grid[0])
    y := len(grid)
    for i := 0; i < y; i++ {
        for j := 0; j < x; j++ {
            if i == 0 && j == 0 {
                continue
            }
            if i == 0 {
                grid[i][j] += grid[i][j-1]
            } else if j == 0 {
                grid[i][j] += grid[i-1][j]
            } else {
                grid[i][j] += min(grid[i-1][j], grid[i][j-1])
            }
        }
    }
    return grid[y-1][x-1]
}
