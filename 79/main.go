package _79

func dfs(board [][]byte, visited [][]bool, i, j int, bytes []byte) bool {
	if len(bytes) == 0 {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) || visited[i][j] || board[i][j] != bytes[0] {
		return false
	} else {
		visited[i][j] = true
		defer func() {
			visited[i][j] = false
		}()
		return dfs(board, visited, i-1, j, bytes[1:]) || dfs(board, visited, i, j+1, bytes[1:]) || dfs(board, visited, i+1, j, bytes[1:]) || dfs(board, visited, i, j-1, bytes[1:])
	}
}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[i]))
	}
	bytes := []byte(word)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if dfs(board, visited, i, j, bytes) {
				return true
			}
		}
	}
	return false
}
