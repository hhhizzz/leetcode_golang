package _541

func reverse(s []byte) {
	half := len(s) >> 1
	for i := 0; i < half; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

func reverseStr(s string, k int) string {
	bytes := []byte(s)
	flag := true
	for i := 0; i < len(s); i += k {
		if flag {
			if i+k < len(s) {
				reverse(bytes[i : i+k])
			} else {
				reverse(bytes[i:])
			}
		}
		flag = !flag
	}
	return string(bytes)
}
