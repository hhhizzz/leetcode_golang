package _54

func spiralOrder(matrix [][]int) []int {
	// 使用方向数组
	// 使用显示器坐标系，左上角为0, 0 x正方向向下，y正方向向右
	direct := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	n := len(matrix)
	if n == 0 {
		return []int{}
	}
	m := len(matrix[0])
	res := make([]int, m*n)
	x, y, d := 0, 0, 0
	visited := make([][]bool, n)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, m)
	}
	for i := 0; i < len(res); i++ {
		res[i] = matrix[x][y]
		visited[x][y] = true
		nextX, nextY := x+direct[d][0], y+direct[d][1]
		if nextX < 0 || nextX >= n || nextY < 0 || nextY >= m || visited[nextX][nextY] {
			d = (d + 1) % 4
			nextX, nextY = x+direct[d][0], y+direct[d][1]
		}
		x, y = nextX, nextY
	}
	return res
}
