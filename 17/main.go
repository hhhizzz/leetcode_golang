package _17

func combine(strs [][]byte, s []byte) [][]byte {
	var res [][]byte
	if len(strs) == 0 {
		for _, a := range s {
			res = append(res, []byte{a})
		}
	} else {
		for _, str := range strs {
			for _, a := range s {
				newStr := make([]byte, len(str))
				copy(newStr, str)
				newStr = append(newStr, a)
				res = append(res, newStr)
			}
		}
	}
	return res
}

func letterCombinations(digits string) []string {
	m := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	var res [][]byte
	for i := 0; i < len(digits); i++ {
		res = combine(res, m[digits[i]])
	}
	var result []string
	for _, s := range res {
		result = append(result, string(s))
	}
	return result
}
