package _463

//广搜1
func islandPerimeter1(grid [][]int) int {
    visited := map[[2]int]bool{}
    stack := [][2]int{}
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == 1 {
                stack = append(stack, [2]int{i, j})
            }
        }
    }
    result := 0
    for len(stack) != 0 {
        current := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if _, ok := visited[[2]int{current[0], current[1]}]; ok {
            continue
        }
        visited[[2]int{current[0], current[1]}] = true
        if current[0] == 0 || grid[current[0]-1][current[1]] == 0 {
            result++
        } else if current[0] != 0 {
            stack = append(stack, [2]int{current[0] - 1, current[1]})
        }
        if current[0] == len(grid)-1 || grid[current[0]+1][current[1]] == 0 {
            result++
        } else if current[0] != len(grid)-1 {
            stack = append(stack, [2]int{current[0] + 1, current[1]})
        }
        if current[1] == 0 || grid[current[0]][current[1]-1] == 0 {
            result++
        } else if current[1] != 0 {
            stack = append(stack, [2]int{current[0], current[1] - 1})
        }
        if current[1] == len(grid[current[0]])-1 || grid[current[0]][current[1]+1] == 0 {
            result++
        } else if current[1] != len(grid[current[0]])-1 {
            stack = append(stack, [2]int{current[0], current[1] + 1})
        }
    }
    return result
}

//广搜2
func islandPerimeter2(grid [][]int) int {
    visited := make([][]bool, len(grid))
    stack := [][2]int{}
    for i := 0; i < len(grid); i++ {
        visited[i] = make([]bool, len(grid[i]))
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == 1 {
                stack = append(stack, [2]int{i, j})
            }
        }
    }
    result := 0
    for len(stack) != 0 {
        current := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if visited[current[0]][current[1]] {
            continue
        }
        visited[current[0]][current[1]] = true
        if current[0] == 0 || grid[current[0]-1][current[1]] == 0 {
            result++
        } else if current[0] != 0 {
            stack = append(stack, [2]int{current[0] - 1, current[1]})
        }
        if current[0] == len(grid)-1 || grid[current[0]+1][current[1]] == 0 {
            result++
        } else if current[0] != len(grid)-1 {
            stack = append(stack, [2]int{current[0] + 1, current[1]})
        }
        if current[1] == 0 || grid[current[0]][current[1]-1] == 0 {
            result++
        } else if current[1] != 0 {
            stack = append(stack, [2]int{current[0], current[1] - 1})
        }
        if current[1] == len(grid[current[0]])-1 || grid[current[0]][current[1]+1] == 0 {
            result++
        } else if current[1] != len(grid[current[0]])-1 {
            stack = append(stack, [2]int{current[0], current[1] + 1})
        }
    }
    return result
}

//直接遍历
func islandPerimeter(grid [][]int) int {
    result := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] != 1 {
                continue
            }
            if i == 0 || grid[i-1][j] == 0 {
                result++
            }
            if i == len(grid)-1 || grid[i+1][j] == 0 {
                result++
            }
            if j == 0 || grid[i][j-1] == 0 {
                result++
            }
            if j == len(grid[i])-1 || grid[i][j+1] == 0 {
                result++
            }
        }
    }
    return result
}
