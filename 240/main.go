package _240

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row := len(matrix) - 1
	col := 0
	for row >= 0 && col < len(matrix[0]) {
		if matrix[row][col] == target {
			return true
		} else if target < matrix[row][col] {
			row--
		} else {
			col++
		}
	}
	return false
}
