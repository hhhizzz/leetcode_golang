package _54

func spiralOrder(matrix [][]int) []int {
	var output []int
	if len(matrix) == 0 {
		return output
	}
	visited := make([][]bool, len(matrix))
	for i := 0; i < len(matrix); i++ {
		visited[i] = make([]bool, len(matrix[i]))
	}
	i := 0
	j := -1
	direction := 0
	RIGHT := 0
	DOWN := 1
	LEFT := 2
	UP := 3
	for {
		if direction == RIGHT {
			if j+1 < len(matrix[i]) && !visited[i][j+1] {
				j++
				output = append(output, matrix[i][j])
				visited[i][j] = true
			} else {
				direction = (direction + 1) % 4
			}
		}
		if direction == DOWN {
			if i+1 < len(matrix) && !visited[i+1][j] {
				i++
				output = append(output, matrix[i][j])
				visited[i][j] = true
			} else {
				direction = (direction + 1) % 4
			}
		}
		if direction == LEFT {
			if j-1 >= 0 && !visited[i][j-1] {
				j--
				output = append(output, matrix[i][j])
				visited[i][j] = true
			} else {
				direction = (direction + 1) % 4
			}
		}
		if direction == UP {
			if i-1 >= 0 && !visited[i-1][j] {
				i--
				output = append(output, matrix[i][j])
				visited[i][j] = true
			} else {
				direction = (direction + 1) % 4
			}
		}
		if (i == 0 || visited[i-1][j]) && (j == 0 || visited[i][j-1]) && (i == (len(matrix)-1) || visited[i+1][j]) && (j == (len(matrix[i])-1) || visited[i][j+1]) {
			break
		}
	}
	return output
}
