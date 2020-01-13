package _1139

func check(grid [][]int, startI, startJ int, edge int) bool {
    //left
    for i := startI; i < startI+edge; i++ {
        if grid[i][startJ] != 1 {
            return false
        }
    }
    //down
    for j := startJ + 1; j < startJ+edge; j++ {
        if grid[startI+edge-1][j] != 1 {
            return false
        }
    }
    //right
    for i := startI + edge - 2; i >= startI; i-- {
        if grid[i][startJ+edge-1] != 1 {
            return false
        }
    }
    //up
    for j := startJ + edge - 2; j > startJ; j-- {
        if grid[startI][j] != 1 {
            return false
        }
    }
    return true
}

func largest1BorderedSquare(grid [][]int) int {
    edgeLength := 1
    result := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            for edge := edgeLength; i+edge <= len(grid) && j+edge <= len(grid[i]); edge++ {
                //fmt.Println(i,j,edge)
                if check(grid, i, j, edge) {
                    result = edge * edge
                    edgeLength = edge
                }
            }
        }
    }
    return result
}
