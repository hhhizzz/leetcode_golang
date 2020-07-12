package _73

func setZeroes(matrix [][]int) {
	row1 := false
	col1 := false
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					row1 = true
				}
				if j == 0 {
					col1 = true
				}
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			for j := 0; j < len(matrix[i]); j++ {
				matrix[i][j] = 0
			}
		}
	}
	for j := 1; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			for i := 0; i < len(matrix); i++ {
				matrix[i][j] = 0
			}
		}
	}
	if row1 {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}
	if col1 {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}
