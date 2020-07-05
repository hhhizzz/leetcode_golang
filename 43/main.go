package _43

func add(num1 string, num2 string) string {
	l1 := len(num1)
	l2 := len(num2)
	if l1 > l2 {
		num1, num2 = num2, num1
		l1, l2 = l2, l1
	}

	var carry byte = 0
	result := make([]byte, l2+1)
	for i1, i2 := l1-1, l2-1; i1 >= 0 || i2 >= 0; i1, i2 = i1-1, i2-1 {
		var current byte = 0
		if i1 < 0 {
			n2 := num2[i2] - '0'
			current = n2 + carry
		} else {
			n1 := num1[i1] - '0'
			n2 := num2[i2] - '0'
			current = n1 + n2 + carry
		}
		if current >= 10 {
			current -= 10
			carry = 1
		} else {
			carry = 0
		}
		result[i2+1] = current + '0'
	}
	if carry != 0 {
		result[0] = '1'
	} else {
		result = result[1:]
	}
	return string(result)
}

func multiply(num1 string, num2 string) string {
	l1 := len(num1)
	l2 := len(num2)

	result := "0"
	for i1 := l1 - 1; i1 >= 0; i1-- {
		n1 := num1[i1] - '0'
		temp := make([]byte, l2+1)
		var carry byte = 0
		for i2 := l2 - 1; i2 >= 0; i2-- {
			n2 := num2[i2] - '0'
			current := n1*n2 + carry
			if current >= 10 {
				carry = current / 10
				current = current % 10
			} else {
				carry = 0
			}
			temp[i2+1] = current + '0'
		}
		if carry != 0 {
			temp[0] = carry + '0'
		} else {
			temp = temp[1:]
		}
		for pos := l1 - 1; pos > i1; pos-- {
			temp = append(temp, '0')
		}
		result = add(string(temp), result)
	}
	pos := 0
	for ; pos < len(result)-1; pos++ {
		if result[pos] != '0' {
			break
		}
	}
	result = result[pos:]
	return result
}

func reverse(num []byte) {
	for i := 0; i < len(num)>>1; i++ {
		num[i], num[len(num)-1-i] = num[len(num)-1-i], num[i]
	}
}

// 方法二，最后处理进位，注意处理前缀0的问题
func multiply2(num1 string, num2 string) string {
	a := []byte(num1)
	b := []byte(num2)

	c := make([]int, len(a)+len(b)-1)

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			c[i+j] += int(a[i]-'0') * int(b[j]-'0')
		}
	}
	var res []byte
	carry := 0
	for i := len(c) - 1; i >= 0; i-- {
		c[i] += carry
		current := c[i]%10 + '0'
		carry = c[i] / 10
		res = append(res, byte(current))
	}
	if carry != 0 {
		res = append(res, byte(carry)+'0')
	}
	reverse(res)
	first := 0
	for first < len(res) && res[first] == '0' {
		first++
	}
	res = res[first:]
	if len(res) == 0 {
		return "0"
	}

	return string(res)
}
