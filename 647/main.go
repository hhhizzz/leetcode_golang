package _647

func countPalindromic(array []byte, i, j int) int {
	count := 0
	for i >= 0 && j < len(array) && array[i] == array[j] {
		count += 1
		//fmt.Println(string(array[i:j+1]))
		i--
		j++
	}
	return count
}

func countSubstrings(s string) int {
	array := []byte(s)
	result := 0
	for i := 0; i < len(array); i++ {
		result += countPalindromic(array, i, i)
		if i+1 < len(array) {
			result += countPalindromic(array, i, i+1)
		}
	}
	return result
}
