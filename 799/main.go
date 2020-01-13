package _799

func champagneTower(poured int, query_row int, query_glass int) float64 {
    cups := make([][]float64, 101)
    for i := 0; i < len(cups); i++ {
        cups[i] = make([]float64, i+1)
    }
    cups[0][0] = float64(poured)
    for i := 0; i < 101; i++ {
        for j := 0; j < i+1; j++ {
            if cups[i][j] > 1.0 {
                overflow := cups[i][j] - 1.0
                cups[i][j] = 1.0
                if i != 100 {
                    cups[i+1][j] += overflow / 2
                    cups[i+1][j+1] += overflow / 2
                }
            }
        }
    }
    return cups[query_row][query_glass]
}
