package _57

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func insert(intervals [][]int, newInterval []int) [][]int {
	var left [][]int
	var right [][]int

	for _, interval := range intervals {
		if interval[1] < newInterval[0] {
			left = append(left, interval)
		} else if interval[0] > newInterval[1] {
			right = append(right, interval)
		} else {
			newInterval[0] = min(newInterval[0], interval[0])
			newInterval[1] = max(newInterval[1], interval[1])
		}
	}
	return append(append(left, newInterval), right...)

}
