package _74

func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 {
        return false
    }
    getNumber := func(pos int) int {
        i := pos / len(matrix[0])
        j := pos - i*len(matrix[0])
        return matrix[i][j]
    }
    left := 0
    right := len(matrix) * len(matrix[0])
    for left < right {
        mid := left + ((right - left) >> 1)
        midNumber := getNumber(mid)
        if midNumber == target {
            return true
        } else if midNumber < target {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return false
}
