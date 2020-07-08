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

//方法2 通过两次对称得到
func rotate2(matrix [][]int) {
	n := len(matrix)
	//  /斜线方向对称
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			matrix[i][j], matrix[n-1-j][n-1-i] = matrix[n-1-j][n-1-i], matrix[i][j]
		}
	}
	// -上下方向对称
	for i := 0; i < n>>1; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j], matrix[n-1-i][j] = matrix[n-1-i][j], matrix[i][j]
		}
	}
}
