package _67

func reverse(bytes []byte) {
	for i := 0; i < len(bytes)>>1; i++ {
		bytes[i], bytes[len(bytes)-1-i] = bytes[len(bytes)-1-i], bytes[i]
	}
}

func addBinary(a string, b string) string {
	bytesA := []byte(a)
	bytesB := []byte(b)

	reverse(bytesA)
	reverse(bytesB)

	if len(bytesB) > len(bytesA) {
		bytesA, bytesB = bytesB, bytesA
	}

	var carry byte = 0
	for i := 0; i < len(bytesA); i++ {
		current := bytesA[i] - '0' + carry
		if i < len(bytesB) {
			current += bytesB[i] - '0'
		}
		carry = current / 2
		bytesA[i] = current%2 + '0'
	}
	if carry != 0 {
		bytesA = append(bytesA, '1')
	}

	reverse(bytesA)
	return string(bytesA)
}
