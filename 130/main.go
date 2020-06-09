package _130

func dfs(board [][]byte, visited [][]bool, i, j int, mark bool) {
	if i == -1 || j == -1 || i == len(board) || j == len(board[i]) {
		return
	}
	if visited[i][j] {
		return
	}
	visited[i][j] = true
	if board[i][j] == 'X' {
		return
	}
	if mark {
		board[i][j] = 'X'
	}
	dfs(board, visited, i-1, j, mark)
	dfs(board, visited, i, j-1, mark)
	dfs(board, visited, i+1, j, mark)
	dfs(board, visited, i, j+1, mark)
}

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[i]))
	}
	for i := 0; i < len(board); i++ {
		dfs(board, visited, i, 0, false)
	}
	for i := 0; i < len(board); i++ {
		dfs(board, visited, i, len(board[i])-1, false)
	}
	for j := 0; j < len(board[0]); j++ {
		dfs(board, visited, 0, j, false)
	}
	for j := 0; j < len(board[0]); j++ {
		dfs(board, visited, len(board)-1, j, false)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			dfs(board, visited, i, j, true)
		}
	}
}
