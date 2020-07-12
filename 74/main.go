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

// 方法2 扫描过程也可以二分，但是估计m和n不算大，直接上速度应该差不多
func searchMatrix2(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		if i == len(matrix)-1 || target < matrix[i+1][0] {
			for j := 0; j < len(matrix[i]); j++ {
				if matrix[i][j] == target {
					return true
				} else if matrix[i][j] > target {
					break
				}
			}
			break
		}
	}
	return false
}
