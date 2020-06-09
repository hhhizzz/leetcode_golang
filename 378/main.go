package _378

func findSmall(matrix [][]int, mid int) (int, bool) {
	count := 0
	contain := false
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] <= mid {
				if matrix[i][j] == mid {
					contain = true
				}
				count++
			} else {
				break
			}
		}
	}
	return count, contain
}

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	left := matrix[0][0]
	right := matrix[n-1][n-1]

	for left < right {
		mid := (left + right) >> 1
		num, contain := findSmall(matrix, mid)
		if num > k || (num == k && !contain) {
			right = mid
		} else if num == k && contain {
			return mid
		} else {
			left = mid + 1
		}
	}
	return left
}
