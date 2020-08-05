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

// 尝试用bfs解决一下
func bfs(board [][]byte, i, j int, b byte) {
	if board[i][j] != 'O' {
		return
	}
	stack := [][]int{{i, j}}
	points := [][]int{}
	for len(stack) != 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		points = append(points, current)
		board[current[0]][current[1]] = b
		if current[0] > 0 && board[current[0]-1][current[1]] == 'O' {
			stack = append(stack, []int{current[0] - 1, current[1]})
		}
		if current[0] < len(board)-1 && board[current[0]+1][current[1]] == 'O' {
			stack = append(stack, []int{current[0] + 1, current[1]})
		}
		if current[1] > 0 && board[current[0]][current[1]-1] == 'O' {
			stack = append(stack, []int{current[0], current[1] - 1})
		}
		if current[1] < len(board[current[0]])-1 && board[current[0]][current[1]+1] == 'O' {
			stack = append(stack, []int{current[0], current[1] + 1})
		}
	}
}

func solve2(board [][]byte) {
	if len(board) == 0 {
		return
	}
	for i := 0; i < len(board); i++ {
		bfs(board, i, 0, 'Y')
		bfs(board, i, len(board[i])-1, 'Y')
	}
	for j := 0; j < len(board[0]); j++ {
		bfs(board, 0, j, 'Y')
		bfs(board, len(board)-1, j, 'Y')
	}
	for i := 1; i < len(board)-1; i++ {
		for j := 1; j < len(board[i])-1; j++ {
			bfs(board, i, j, 'X')
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'Y' {
				board[i][j] = 'O'
			}
		}
	}
}
