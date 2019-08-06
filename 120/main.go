package _120

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func minimumTotal(triangle [][]int) int {
    button := len(triangle) - 1
    length := len(triangle[button]) - 1
    for row := button - 1; row >= 0; row-- {
        for col := 0; col < length; col++ {
            triangle[row][col] = triangle[row][col] + min(triangle[row+1][col], triangle[row+1][col+1])
        }
        length--
    }
    return triangle[0][0]
}
