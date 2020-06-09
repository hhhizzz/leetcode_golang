package _394

func decodeString(s string) string {
	var result string
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c <= '9' && c >= '0' {
			times := 0
			for s[i] != '[' {
				times *= 10
				times += int(s[i] - '0')
				i++
			}
			leftBucket := 1
			j := i
			for ; leftBucket != 0; j++ {
				if s[j] == '[' {
					leftBucket++
				} else if s[j] == ']' {
					leftBucket--
				}
			}
			unit := decodeString(s[i+1 : j])
			for time := 1; time <= times; time++ {
				result += unit
			}
			i = j
		} else {
			result += string(c)
		}
	}
	return result
}
