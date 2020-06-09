package _763

func partitionLabels(S string) []int {
	last := map[uint8]int{}
	for i := 0; i < len(S); i++ {
		last[S[i]] = i
	}
	var result []int
	for i := 0; i < len(S); i++ {
		left := i
		right := last[S[i]]
		for j := left + 1; j < right; j++ {
			if last[S[j]] > right {
				right = last[S[j]]
			}
		}
		result = append(result, right-left+1)
		i = right
	}
	return result
}
