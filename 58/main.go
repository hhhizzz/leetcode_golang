package _58

func lengthOfLastWord(s string) int {
	bytes := []byte(s)
	lastC := -1
	for i := len(bytes) - 1; i >= 0; i-- {
		if bytes[i] != ' ' && lastC == -1 {
			lastC = i
		} else if bytes[i] == ' ' && lastC != -1 {
			return lastC - i
		}
	}
	if lastC == -1 {
		return 0
	}
	return lastC + 1
}
