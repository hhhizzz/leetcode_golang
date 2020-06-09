package _796

func equal(s1, s2 []byte) bool {
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func rotateString(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	if A == B {
		return true
	}
	s1 := []byte(A)
	s2 := []byte(B)

	for i := 0; i < len(s1); i++ {
		if equal(s1, s2) {
			return true
		}
		s1 = append(s1[1:], s1[0])
	}
	return false
}
