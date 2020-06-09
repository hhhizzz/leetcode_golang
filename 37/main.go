package _37

func isValid(board *[][]byte, row, col int, number byte) bool {
	for i := 0; i < 9; i++ {
		if (*board)[i][col] == number {
			return false
		}
	}
	for j := 0; j < 9; j++ {
		if (*board)[row][j] == number {
			return false
		}
	}
	x := row - row%3
	y := col - col%3
	for i := x; i < x+3; i++ {
		for j := y; j < y+3; j++ {
			if (*board)[i][j] == number {
				return false
			}
		}
	}
	return true
}
func solve(board *[][]byte, i, j int) bool {
	if (*board)[i][j] != '.' {
		if i == 8 && j == 8 {
			return true
		}
		if j < 8 {
			return solve(board, i, j+1)
		} else {
			return solve(board, i+1, 0)
		}
	}
	var number byte
	for number = 1; number <= 9; number++ {
		if isValid(board, i, j, number+'0') {
			(*board)[i][j] = number + '0'
			if i == 8 && j == 8 {
				return true
			}
			var result bool
			if j < 8 {
				result = solve(board, i, j+1)
			} else {
				result = solve(board, i+1, 0)
			}
			if result {
				return true
			}
		}
	}
	(*board)[i][j] = '.'
	return false
}

func solveSudoku(board [][]byte) {
	solve(&board, 0, 0)
}
