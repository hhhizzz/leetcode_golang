package _20

func findLeft(c byte) byte {
	if c == ')' {
		return '('
	}
	if c == ']' {
		return '['
	}
	if c == '}' {
		return '{'
	}
	return ' '
}

func isLeft(c byte) bool {
	return c == '(' || c == '[' || c == '{'
}

func isValid(s string) bool {
	var stack []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		if isLeft(c) {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != findLeft(c) {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}
