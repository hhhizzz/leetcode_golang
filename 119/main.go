package _119

func getRow(rowIndex int) []int {
	res := []int{1}
	for i := 1; i <= rowIndex; i++ {
		next := []int{1}
		for i := 1; i < len(res); i++ {
			next = append(next, res[i]+res[i-1])
		}
		next = append(next, 1)
		res = next
	}
	return res
}
