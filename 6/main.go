package _6

//找规律题，每行都是等差数列，差为2n-2
/*
1      7        13
2    6 8     12 14
3  5   9  11    15
4      10       16
*/
func convert(s string, numRows int) string {
	//注意边缘情况，n为1的时候，规律不成立，直接返回原字符串
	if numRows <= 1 {
		return s
	}
	length := numRows*2 - 2
	var res []byte
	for i := 0; i < numRows; i++ {
		if i == 0 || i == numRows-1 {
			for j := i; j < len(s); j += length {
				res = append(res, s[j])
			}
		} else {
			for j, k := i, length-i; j < len(s) || k < len(s); {
				if j < len(s) {
					res = append(res, s[j])
				}
				if k < len(s) {
					res = append(res, s[k])
				}
				j += length
				k += length
			}
		}
	}
	return string(res)
}
