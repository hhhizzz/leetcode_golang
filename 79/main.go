package _79

func search(board *[][]byte, visited *[][]bool, word *[]byte, i, j, pos int) bool {
	if i < 0 || j < 0 || i >= len(*board) || j >= len((*board)[i]) {
		return false
	}

	if (*visited)[i][j] {
		return false
	}

	if (*word)[pos] == (*board)[i][j] {
		if pos == len(*word)-1 {
			return true
		} else {
			(*visited)[i][j] = true
			if j != len((*board)[i]) {
				if search(board, visited, word, i-1, j, pos+1) {
					return true
				}
				if search(board, visited, word, i, j+1, pos+1) {
					return true
				}
				if search(board, visited, word, i+1, j, pos+1) {
					return true
				}
				if search(board, visited, word, i, j-1, pos+1) {
					return true
				}
			} else {
				j = 0
				i += 1
				if search(board, visited, word, i-1, j, pos+1) {
					return true
				}
				if search(board, visited, word, i, j+1, pos+1) {
					return true
				}
				if search(board, visited, word, i+1, j, pos+1) {
					return true
				}
			}
			(*visited)[i][j] = false
		}
	}
	return false
}

func exist(board [][]byte, word string) bool {
	wordArray := []byte(word)
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[i]))
	}
	for i := range board {
		for j := range board[i] {
			if search(&board, &visited, &wordArray, i, j, 0) {
				return true
			}
		}
	}
	return false
}
