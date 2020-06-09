package _873

func lenLongestFibSubseq(A []int) int {
	result := 0
	aMap := map[int]int{}
	for i := 0; i < len(A); i++ {
		aMap[A[i]] = i
	}
	for i := 1; i < len(A); i++ {
		for j := i - 1; j >= 0; j-- {
			currentLength := 2
			current := i
			next := j
			for {
				nextNumber := A[current] - A[next]
				if nextNumber >= A[next] {
					break
				}
				nextNext, ok := aMap[nextNumber]
				if !ok {
					break
				} else {
					currentLength += 1
					if currentLength > result {
						result = currentLength
					}
					current, next = next, nextNext
				}
			}
		}
	}
	return result
}
