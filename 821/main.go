package _821

func shortestToChar(S string, C byte) []int {
	array := []byte(S)
	result := make([]int, len(S))
	var pos []int
	for i := 0; i < len(S); i++ {
		if array[i] == C {
			pos = append(pos, i)
		} else {
			result[i] = 1
		}
	}
	for i := 0; i < len(pos); i++ {
		if i == 0 {
			for j := pos[i] - 1; j >= 0; j-- {
				result[j] = pos[i] - j
			}

		}
		if i == len(pos)-1 {
			for j := pos[i] + 1; j < len(S); j++ {
				result[j] = j - pos[i]
			}
		} else {
			left := pos[i] + 1
			right := pos[i+1] - 1
			dis := 1
			for left <= right {
				result[left] = dis
				result[right] = dis
				dis++
				left++
				right--
			}
		}
	}
	return result
}
