package _8

import "math"

func myAtoi(str string) int {
	s := []byte(str)
	for len(s) > 0 && s[0] == ' ' {
		s = s[1:]
	}
	if len(s) == 0 {
		return 0
	}
	minus := 1
	if s[0] == '-' {
		minus = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}
	res := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			x := int(c - '0')
			//难点在于用int32处理越界问题
			if minus > 0 && res > (math.MaxInt32-x)/10 {
				return math.MaxInt32
			}
			if minus < 0 && -res < (math.MinInt32+x)/10 {
				return math.MinInt32
			}
			// 注意特别情况，-2147483648 这个值的负数没有越界，但是正数越界了，res不能正确表达这个值，需要专门检查
			if -res*10-x == math.MinInt32 {
				return math.MinInt32
			}
			res = res*10 + x
		} else {
			break
		}
	}
	return res * minus
}
