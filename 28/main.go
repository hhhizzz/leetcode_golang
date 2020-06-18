package _28

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	next := make([]int, len(needle))
	next[0] = -1
	for i, j := 0, -1; i < len(next)-1; {
		if j == -1 || needle[i] == needle[j] {
			i++
			j++
			if needle[i] == needle[j] {
				next[i] = next[j]
			} else {
				next[i] = j
			}
		} else {
			j = next[j]
		}
	}

	for i, j := 0, 0; i < len(haystack); {
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
			if j == len(needle) {
				return i - j
			}
		} else {
			j = next[j]
		}
	}
	return -1
}
