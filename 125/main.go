package _125

func word(c byte) bool {
	if (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
		return true
	}
	return false
}

func lower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c - 'A' + 'a'
	}
	return c
}

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if !word(s[left]) {
			left++
		} else if !word(s[right]) {
			right--
		} else if lower(s[left]) != lower(s[right]) {
			return false
		} else {
			left++
			right--
		}
	}
	return true
}
