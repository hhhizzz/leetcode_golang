package _12

func intToRoman(num int) string {
	var res []byte
	for num != 0 {
		if num >= 1000 {
			res = append(res, 'M')
			num -= 1000
		} else if num >= 500 {
			if num >= 900 {
				res = append(res, 'C')
				res = append(res, 'M')
				num -= 900
			} else {
				res = append(res, 'D')
				num -= 500
			}

		} else if num >= 100 {
			if num >= 400 {
				res = append(res, 'C')
				res = append(res, 'D')
				num -= 400
			} else {
				res = append(res, 'C')
				num -= 100
			}

		} else if num >= 50 {
			if num >= 90 {
				res = append(res, 'X')
				res = append(res, 'C')
				num -= 90
			} else {
				res = append(res, 'L')
				num -= 50
			}
		} else if num >= 10 {
			if num >= 40 {
				res = append(res, 'X')
				res = append(res, 'L')
				num -= 40
			} else {
				res = append(res, 'X')
				num -= 10
			}
		} else if num >= 5 {
			if num >= 9 {
				res = append(res, 'I')
				res = append(res, 'X')
				num -= 9
			} else {
				res = append(res, 'V')
				num -= 5
			}
		} else {
			if num >= 4 {
				res = append(res, 'I')
				res = append(res, 'V')
				num -= 4
			} else {
				res = append(res, 'I')
				num -= 1
			}
		}
	}
	return string(res)
}
