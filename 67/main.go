package _67

func addBinary(a string, b string) string {
	var arrayA []byte
	var arrayB []byte
	for i := len(a) - 1; i >= 0; i-- {
		arrayA = append(arrayA, a[i])
	}
	for i := len(b) - 1; i >= 0; i-- {
		arrayB = append(arrayB, b[i])
	}
	if len(a) < len(b) {
		arrayA, arrayB = arrayB, arrayA
	}
	var carry byte = 0
	for i := 0; i < len(arrayA); i++ {
		if i < len(arrayB) {
			current := carry + (arrayA[i] - '0') + (arrayB[i] - '0')
			if current >= 2 {
				carry = 1
			} else {
				carry = 0
			}
			arrayA[i] = current%2 + '0'
		} else {
			current := carry + (arrayA[i] - '0')
			if current >= 2 {
				carry = 1
			} else {
				carry = 0
			}
			arrayA[i] = current%2 + '0'
		}
	}
	if carry != 0 {
		arrayA = append(arrayA, '1')
	}
	for i := 0; i < len(arrayA)/2; i++ {
		arrayA[i], arrayA[len(arrayA)-i-1] = arrayA[len(arrayA)-i-1], arrayA[i]
	}
	return string(arrayA)
}
