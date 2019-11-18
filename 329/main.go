package _329

type node struct {
    num   int
    level int
}

//方法一 使用拓扑排序
func longestIncreasingPath(matrix [][]int) int {
    if len(matrix) == 0 {
        return 0
    }
    g := make([][]int, len(matrix)*len(matrix[0]))
    inDegree := make([]int, len(matrix)*len(matrix[0]))
    var stack []node
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            num := i*len(matrix[i]) + j
            //left
            if j != 0 && matrix[i][j-1] < matrix[i][j] {
                g[num] = append(g[num], num-1)
                inDegree[num-1]++
            }
            //up
            if i != 0 && matrix[i-1][j] < matrix[i][j] {
                g[num] = append(g[num], num-len(matrix[i]))
                inDegree[num-len(matrix[i])]++
            }
            //right
            if j < len(matrix[i])-1 && matrix[i][j+1] < matrix[i][j] {
                g[num] = append(g[num], num+1)
                inDegree[num+1]++
            }
            //button
            if i < len(matrix)-1 && matrix[i+1][j] < matrix[i][j] {
                g[num] = append(g[num], num+len(matrix[i]))
                inDegree[num+len(matrix[i])]++
            }
        }
    }
    for i := 0; i < len(inDegree); i++ {
        if inDegree[i] == 0 {
            stack = append(stack, node{level: 1, num: i})
        }
    }
    result := 0
    for len(stack) != 0 {
        current := stack[0]
        stack = stack[1:]
        next := g[current.num]
        for _, n := range next {
            inDegree[n]--
            if inDegree[n] == 0 {
                stack = append(stack, node{level: current.level + 1, num: n})
            }
        }
        if current.level > result {
            result = current.level
        }
    }
    return result
}
