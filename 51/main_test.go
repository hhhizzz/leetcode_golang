package _51

import "testing"

func TestSolveNQueens(t *testing.T) {
	queens := solveNQueens(8)
	println(len(queens))
	for _, queen := range queens {
		for _, q := range queen {
			println(q)
		}
		println()
	}
}
