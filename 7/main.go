package _7

import "math"

func reverse(x int) int {
	result := 0
	for x != 0 {
		current := x % 10
		result += current

		x /= 10
		if x != 0 {
			if result > math.MaxInt32/10 {
				return 0
			}
			if result < math.MinInt32/10 {
				return 0
			}
			result *= 10
		}
	}
	return result
}
