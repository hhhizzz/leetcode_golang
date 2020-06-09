package _739

//单调栈解法
func dailyTemperatures(T []int) []int {
	var s []int
	result := make([]int, len(T))

	for i := 0; i < len(T); i++ {
		if len(s) == 0 || T[i] <= T[s[len(s)-1]] {
			s = append(s, i)
		} else {
			for len(s) > 0 && T[s[len(s)-1]] < T[i] {
				top := s[len(s)-1]
				s = s[:len(s)-1]
				result[top] = i - top
			}
			s = append(s, i)
		}
	}
	return result
}
