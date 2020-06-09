package _79

import "testing"

func TestExist(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	println(exist(board, "ABCCED"))
	println(exist(board, "SEE"))
	println(exist(board, "ABCE"))
	println(exist([][]byte{{'a'}}, "ab"))
	println(exist([][]byte{{'a', 'b'}}, "ba"))
}
