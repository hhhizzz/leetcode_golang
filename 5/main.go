package _5

import "strings"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func longestPalindrome(s string) string {
	var pos, maxLeft, maxRight, maxLength int
	maxLeft = 0
	maxRight = 0
	maxLength = 1
	var sSharp = "#"
	for _, c := range s {
		sSharp += string(c) + "#"
	}

	var length = len(sSharp)

	RL := make([]int, length)
	for i := 0; i < length; i++ {
		if i < maxRight {
			j := 2*pos - i
			RL[i] = min(RL[j], maxRight-i)
		}
		var left = i - RL[i]
		var right = i + RL[i]
		for left >= 0 && right < length && sSharp[left] == sSharp[right] {
			left--
			right++
		}
		if left == right {
			continue
		} else {
			left++
			right--
		}
		currentLength := right - left + 1
		RL[i] = right - i + 1
		if currentLength > maxLength {
			pos = i
			maxLeft = left
			maxRight = right
			maxLength = currentLength
		}
	}
	var result string
	result = strings.Replace(sSharp[maxLeft:maxRight], "#", "", -1)
	return result
}
