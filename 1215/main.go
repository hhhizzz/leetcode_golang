package _215

func countSteppingNumbers(low int, high int) []int {
	var queue []int
	var result []int

	for i := 0; i < 10 && i <= high; i++ {
		result = append(result, i)
		if i != 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) != 0 {
		first := queue[0]
		queue = queue[1:]

		for i := 0; i < 10; i++ {
			if first%10-i == 1 || i-first%10 == 1 {
				next := first*10 + i
				if next <= high {
					result = append(result, next)
					queue = append(queue, next)
				}
			}
		}
	}
	left := 0
	right := len(result)
	for left < right {
		mid := (left + right) >> 1
		if result[mid] > low {
			right = mid
		} else if result[mid] < low {
			left = mid + 1
		} else {
			left = mid
			break
		}
	}
	return result[left:]
}
