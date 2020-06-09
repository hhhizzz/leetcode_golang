package _171

func titleToNumber(s string) int {
	result := 1
	for i, c := range s {
		if i != 0 {
			result *= 26
			result += 1
		}
		result += int(c) - 'A'
	}
	return result
}
