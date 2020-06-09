package _5

func palindrome(array []byte, l, r int) []byte {
	for l >= 0 && r < len(array) && array[l] == array[r] {
		l--
		r++
	}
	return array[l+1 : r]
}

func longestPalindrome(s string) string {
	var result []byte
	array := []byte(s)
	for i := 0; i < len(s); i++ {
		current := palindrome(array, i, i)
		if len(current) > len(result) {
			result = current
		}
		current = palindrome(array, i, i+1)
		if len(current) > len(result) {
			result = current
		}
	}
	return string(result)
}
