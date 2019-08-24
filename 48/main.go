package _48

var n int

func nextPos(x, y int) (int, int) {
    return n - 1 - y, x
}

func rotate(matrix [][]int) {
    n = len(matrix)
    for i := 0; i < n>>1; i++ {
        for j := i; j < n-i-1; j++ {
            x, y := j, i
            next := matrix[y][x]
            for times := 0; times < 4; times++ {
                nextX, nextY := nextPos(x, y)
                temp := matrix[nextY][nextX]
                matrix[nextY][nextX] = next
                next = temp
                x, y = nextX, nextY
            }
        }
    }
}
