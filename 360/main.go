package _360

func sortTransformedArray(nums []int, a int, b int, c int) []int {
	var qua []int
	for _, num := range nums {
		qua = append(qua, a*num*num+b*num+c)
	}
	var result []int
	//开口向上
	if a > 0 {
		for i := 1; i <= len(qua); i++ {
			if i == len(qua) || qua[i] > qua[i-1] {
				if i == 1 {
					result = qua
				} else {
					left := i - 1
					right := i
					for left != -1 || right != len(qua) {
						if (right != len(qua)) && (left == -1 || qua[left] >= qua[right]) {
							result = append(result, qua[right])
							right++
						}
						if (left != -1) && (right == len(qua) || qua[left] <= qua[right]) {
							result = append(result, qua[left])
							left--
						}
					}
				}
				break
			}
		}
	} else if a < 0 {
		//开口向下
		for i := 1; i <= len(qua); i++ {
			if i == len(qua) {
				result = qua
				break
			} else if qua[i-1] > qua[i] {
				left := 0
				right := len(qua) - 1
				for left <= right {
					if left <= right && qua[left] >= qua[right] {
						result = append(result, qua[right])
						right--
					}
					if left <= right && qua[left] <= qua[right] {
						result = append(result, qua[left])
						left++
					}
				}
				break
			}
		}
	} else {
		if b > 0 {
			result = qua
		} else {
			for i := 0; i < len(qua); i++ {
				result = append(result, qua[len(qua)-1-i])
			}
		}
	}
	return result
}
