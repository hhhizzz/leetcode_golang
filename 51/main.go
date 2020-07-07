package _51

func dfs(board [][]byte, i int, row []bool, slash []bool, bSlash []bool, result *[][]string) {
	n := len(board)
	if i == len(board) {
		current := make([]string, n)
		for x := 0; x < n; x++ {
			current[x] = string(board[x])
		}
		*result = append(*result, current)
	}
	for j := 0; j < n; j++ {
		if !row[j] && !slash[i+j] && !bSlash[i-j+n-1] {
			board[i][j] = 'Q'
			row[j] = true
			slash[i+j] = true
			bSlash[i-j+n-1] = true
			dfs(board, i+1, row, slash, bSlash, result)
			board[i][j] = '.'
			row[j] = false
			slash[i+j] = false
			bSlash[i-j+n-1] = false
		}
	}
}

func solveNQueens(n int) [][]string {
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}
	row := make([]bool, n)
	slash := make([]bool, 2*n-1)
	bSlash := make([]bool, 2*n-1)
	var result [][]string
	dfs(board, 0, row, slash, bSlash, &result)
	return result
}
