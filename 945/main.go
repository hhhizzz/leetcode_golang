package _945

func minIncrementForUnique(A []int) int {
	//注意这个值应该是最大值的两倍
	count := [80001]int{}
	for i := 0; i < len(A); i++ {
		count[A[i]]++
	}
	result := 0
	taken := 0
	for i := 0; i <= 80000; i++ {
		if count[i] > 1 {
			duplicate := count[i] - 1
			taken += duplicate
			result -= i * duplicate
		} else if count[i] == 0 && taken > 0 {
			result += i
			taken -= 1
		}
	}
	return result
}
