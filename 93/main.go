package _93

func toInt(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res *= 10
		res += int(s[i] - '0')
	}
	return res
}

func helper(s string, current string, number int, res *[]string) {
	if number == 4 {
		if len(s) == 0 {
			*res = append(*res, current[:len(current)-1])
		} else {
			return
		}
	}
	for i := 1; i <= len(s); i++ {
		if i > 1 && s[0] == '0' {
			break
		}
		num := toInt(s[:i])
		if num <= 255 {
			helper(s[i:], current+s[:i]+".", number+1, res)
		} else {
			break
		}
	}
}

func restoreIpAddresses(s string) []string {
	var res []string
	helper(s, "", 0, &res)
	return res
}
