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

// 第一次写
func solveSudoku(board [][]byte) {
	solve(&board, 0, 0)
}

func toBoxNumber(x, y int) int {
	return x/3*3 + y/3
}

func nextEmpty(board [][]byte) (int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				return i, j
			}
		}
	}
	return -1, -1
}

func helper(board [][]byte, row [9][10]bool, col [9][10]bool, box [9][10]bool, i, j int) bool {
	for number := 1; number <= 9; number++ {
		if !col[j][number] && !row[i][number] && !box[toBoxNumber(i, j)][number] {
			board[i][j] = byte(number) + '0'
			row[i][number] = true
			col[j][number] = true
			box[toBoxNumber(i, j)][number] = true

			nextI, nextJ := nextEmpty(board)
			if nextI == -1 && nextJ == -1 {
				return true
			} else {
				res := helper(board, row, col, box, nextI, nextJ)
				if res {
					return res
				} else {
					board[i][j] = '.'
					row[i][number] = false
					col[j][number] = false
					box[toBoxNumber(i, j)][number] = false
				}
			}
		}
	}
	return false
}

// 第二次尝试
func solveSudoku2(board [][]byte) {
	var row [9][10]bool
	var col [9][10]bool
	var box [9][10]bool
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				number := board[i][j] - '0'
				row[i][number] = true
				col[j][number] = true
				box[toBoxNumber(i, j)][number] = true
			}
		}
	}
	i, j := nextEmpty(board)
	helper(board, row, col, box, i, j)
}
