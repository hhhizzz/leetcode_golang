package _73

func setZeroes(matrix [][]int) {
    row := map[int]bool{}
    col := map[int]bool{}
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            if matrix[i][j] == 0 {
                col[j] = true
                row[i] = true
            }
        }
    }
    for r := range row {
        for j := 0; j < len(matrix[r]); j++ {
            matrix[r][j] = 0
        }
    }
    for c := range col {
        for i := 0; i < len(matrix); i++ {
            matrix[i][c] = 0
        }
    }
}
