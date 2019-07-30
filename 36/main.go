package _36

func isValidSudoku(board [][]byte) bool {
    col := make([]bool, 9)
    row := make([]bool, 9)
    block := make([]bool, 9)
    //check if row is valid
    for i := 0; i < 9; i++ {
        row = make([]bool, 9)
        for j := 0; j < 9; j++ {
            numberByte := board[i][j]
            var number byte
            if numberByte == '.' {
                continue
            } else {
                number = numberByte - '0'
            }
            if exist := row[number-1]; exist {
                return false
            } else {
                row[number-1] = true
            }
        }
    }
    //check if col is valid
    for j := 0; j < 9; j++ {
        col = make([]bool, 9)
        for i := 0; i < 9; i++ {
            numberByte := board[i][j]
            var number byte
            if numberByte == '.' {
                continue
            } else {
                number = numberByte - '0'
            }
            if exist := col[number-1]; exist {
                return false
            } else {
                col[number-1] = true
            }
        }
    }
    //check if block is valid
    for x := 0; x < 9; x += 3 {
        for y := 0; y < 9; y += 3 {
            block = make([]bool, 9)
            for i := 0; i < 3; i++ {
                for j := 0; j < 3; j++ {
                    numberByte := board[x+i][y+j]
                    var number byte
                    if numberByte == '.' {
                        continue
                    } else {
                        number = numberByte - '0'
                    }
                    if exist := block[number-1]; exist {
                        return false
                    } else {
                        block[number-1] = true
                    }
                }
            }
        }
    }
    return true
}
