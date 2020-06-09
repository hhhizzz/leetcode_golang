package _51

var num int
var col map[int]bool
var pie map[int]bool
var na map[int]bool
var result [][]string

func dfs(row int, current *[]string) {
	for i := 0; i < num; i++ {
		_, colOK := col[i]
		_, pieOK := pie[i+row]
		_, naOK := na[i-row]
		if !colOK && !pieOK && !naOK {
			col[i] = true
			pie[i+row] = true
			na[i-row] = true
			newCol := ""
			for j := 0; j < num; j++ {
				if j != i {
					newCol += "."
				} else {
					newCol += "Q"
				}
			}
			*current = append(*current, newCol)
			if row == num-1 {
				newResult := make([]string, num)
				copy(newResult, *current)
				result = append(result, newResult)
			} else {
				dfs(row+1, current)
			}
			delete(col, i)
			delete(pie, i+row)
			delete(na, i-row)
			*current = (*current)[:row]
		}
	}
}

func totalNQueens(n int) int {
	return len(solveNQueens(n))
}

func solveNQueens(n int) [][]string {
	col = make(map[int]bool)
	pie = make(map[int]bool)
	na = make(map[int]bool)
	result = [][]string{}
	num = n
	dfs(0, &[]string{})
	return result
}
