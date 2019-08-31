package _59

const RIGHT = 0
const DONW = 1
const LEFT = 2
const UP = 3

func generateMatrix(n int) [][]int {
    output := make([][]int, n)
    for i := 0; i < n; i++ {
        output[i] = make([]int, n)
    }
    direction := RIGHT
    n2 := n * n
    i := 0
    j := -1
    for num := 1; num <= n2; {
        if direction == RIGHT {
            for j+1 < len(output[i]) && output[i][j+1] == 0 {
                j++
                output[i][j] = num
                num++
            }
            direction = (direction + 1) % 4
        }
        if direction == DONW {
            for i+1 < len(output) && output[i+1][j] == 0 {
                i++
                output[i][j] = num
                num++
            }
            direction = (direction + 1) % 4
        }
        if direction == LEFT {
            for j-1 >= 0 && output[i][j-1] == 0 {
                j--
                output[i][j] = num
                num++
            }
            direction = (direction + 1) % 4
        }
        if direction == UP {
            for i-1 >=0 && output[i-1][j] == 0 {
                i--
                output[i][j] = num
                num++
            }
            direction = (direction + 1) % 4
        }

    }
    return output
}
