package _221

var length [][]int
var height int

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func checkSquare(i, j int) int {
	if length[i][j] <= 1 {
		return length[i][j]
	}
	maxLength := length[i][j]
	currentLength := 1

	for newI := i + 1; newI < i+maxLength && newI < height; newI++ {
		maxLength = min(maxLength, length[newI][j])
		currentLength = min(maxLength, currentLength+1)
	}
	return currentLength * currentLength
}

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	height = len(matrix)

	length = make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		length[i] = make([]int, len(matrix[i]))
		current := 1
		for j := len(matrix[i]) - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				length[i][j] = current
				current++
			} else {
				current = 1
			}
		}
	}
	result := 0
	for i := 0; i < len(length); i++ {
		for j := 0; j < len(length[i]); j++ {
			result = max(result, checkSquare(i, j))
		}
	}
	return result
}
