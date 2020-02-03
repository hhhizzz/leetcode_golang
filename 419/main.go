package _419

func countBattleships(board [][]byte) int {
    result := 0
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            if board[i][j] == 'X' {
                if i != len(board)-1 && board[i+1][j] == 'X' {
                    continue
                } else {
                    result++
                    for j+1 < len(board[i]) && board[i][j+1] == 'X' {
                        j++
                    }
                }
            }

        }
    }
    return result
}
