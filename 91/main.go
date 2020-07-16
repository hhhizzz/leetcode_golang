package _91

func numDecodings(s string) int {
	l, m := 1, 0
	res := 0
	if len(s) > 0 && s[0] != '0' {
		res = 1
		m = 1
	}
	for i := 1; i < len(s); i++ {
		n := s[i] - '0'
		nn := n + 10*(s[i-1]-'0')
		if (nn >= 21 && nn <= 26) || (nn >= 11 && nn <= 19) {
			res = l + m
		} else if nn == 20 || nn == 10 {
			res = l
		} else if n == 0 {
			return 0
		} else {
			res = m
		}
		l, m = m, res
	}
	return res
}
