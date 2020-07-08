package _59

const RIGHT = 0
const DONW = 1
const LEFT = 2
const UP = 3

func generateMatrix(n int) [][]int {
	output := make([][]int, n)
	for i := 0; i < n; i++ {
		output[i] = make([]int, n)
	}
	direction := RIGHT
	n2 := n * n
	i := 0
	j := -1
	for num := 1; num <= n2; {
		if direction == RIGHT {
			for j+1 < len(output[i]) && output[i][j+1] == 0 {
				j++
				output[i][j] = num
				num++
			}
			direction = (direction + 1) % 4
		}
		if direction == DONW {
			for i+1 < len(output) && output[i+1][j] == 0 {
				i++
				output[i][j] = num
				num++
			}
			direction = (direction + 1) % 4
		}
		if direction == LEFT {
			for j-1 >= 0 && output[i][j-1] == 0 {
				j--
				output[i][j] = num
				num++
			}
			direction = (direction + 1) % 4
		}
		if direction == UP {
			for i-1 >= 0 && output[i-1][j] == 0 {
				i--
				output[i][j] = num
				num++
			}
			direction = (direction + 1) % 4
		}

	}
	return output
}

// 方向数组方案
func generateMatrix2(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	direct := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	x, y := 0, 0
	point := 0
	for i := 1; i <= n*n; i++ {
		res[x][y] = i
		nextX, nextY := x+direct[point][0], y+direct[point][1]
		if nextX < 0 || nextX >= n || nextY < 0 || nextY >= n || res[nextX][nextY] != 0 {
			point = (point + 1) % len(direct)
			nextX, nextY = x+direct[point][0], y+direct[point][1]
		}
		x, y = nextX, nextY
	}
	return res
}
