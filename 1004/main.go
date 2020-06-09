package _1004

func longestOnes(A []int, K int) int {
	if K > len(A) {
		return len(A)
	}
	left := 0
	right := 0
	numZero := 0
	result := 0
	for right < len(A) {
		if A[right] == 0 {
			numZero++
		}
		for numZero > K {
			if A[left] == 0 {
				numZero--
			}
			left++
		}
		if right-left+1 > result {
			result = right - left + 1
		}
		right++

		//fmt.Printf("window: %v, numZero: %d\n",A[left:right],numZero)
	}
	return result
}
