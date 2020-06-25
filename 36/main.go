package _36

func isValid(array []byte) bool {
	m := make([]bool, 10)
	for i := 0; i < 9; i++ {
		if array[i] == '.' {
			continue
		} else {
			number := array[i] - '0'
			if m[number] {
				return false
			}
			m[number] = true
		}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if !isValid(board[i]) {
			return false
		}
	}
	for j := 0; j < 9; j++ {
		temp := make([]byte, 9)
		for i := 0; i < 9; i++ {
			temp[i] = board[i][j]
		}
		if !isValid(temp) {
			return false
		}
	}
	for k := 0; k < 9; k++ {
		row := (k / 3) * 3
		col := (k % 3) * 3
		temp := make([]byte, 9)
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				temp[i*3+j] = board[row+i][col+j]
			}
		}
		if !isValid(temp) {
			return false
		}
	}
	return true
}
