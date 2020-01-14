package _348

type TicTacToe struct {
    table [][]int
}

/** Initialize your data structure here. */
func Constructor(n int) TicTacToe {
    table := make([][]int, n)
    for i := 0; i < n; i++ {
        table[i] = make([]int, n)
    }
    return TicTacToe{table: table}
}

/** Player {player} makes a move at ({row}, {col}).
  @param row The row of the board.
  @param col The column of the board.
  @param player The player, can be either 1 or 2.
  @return The current winning condition, can be either:
          0: No one wins.
          1: Player 1 wins.
          2: Player 2 wins. */
func (this *TicTacToe) Move(row int, col int, player int) int {
    this.table[row][col] = player
    // fmt.Println(this.table)
    // fmt.Println()

    win := true
    for i := 0; i < len(this.table); i++ {
        if this.table[i][col] != player {
            win = false
            break
        }
    }
    if win {
        return player
    }

    win = true
    for j := 0; j < len(this.table); j++ {
        if this.table[row][j] != player {
            win = false
            break
        }
    }
    if win {
        return player
    }

    win = true
    if row != col {
        win = false
    } else {
        for i := 0; i < len(this.table); i++ {
            if this.table[i][i] != player {
                win = false
                break
            }
        }
    }
    if win {
        return player
    }

    win = true
    if row != len(this.table)-1-col {
        win = false
    } else {
        for i := 0; i < len(this.table); i++ {
            if this.table[i][len(this.table)-1-i] != player {
                win = false
                break
            }
        }
    }
    if win {
        return player
    }
    return 0
}

/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */
