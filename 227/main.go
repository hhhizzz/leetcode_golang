package _227

func calculate(str string) int {
	var post []interface{}
	in := []byte(str)

	var s []byte
	for i := 0; i < len(in); i++ {
		if in[i] == ' ' || in[i] == '\t' {
			continue
		} else if in[i] >= '0' && in[i] <= '9' {
			number := 0
			for i < len(in) && in[i] >= '0' && in[i] <= '9' {
				number *= 10
				number += int(in[i] - '0')
				i++
			}
			post = append(post, number)
			i--
		} else {
			if len(s) == 0 {
				s = append(s, in[i])
			} else {
				if in[i] == '+' || in[i] == '-' {
					for len(s) > 0 {
						top := s[len(s)-1]
						post = append(post, top)
						s = s[:len(s)-1]
					}
					s = append(s, in[i])
				} else if in[i] == '*' || in[i] == '/' {
					for len(s) > 0 {
						top := s[len(s)-1]
						if top == '*' || top == '/' {
							post = append(post, top)
							s = s[:len(s)-1]
						} else {
							break
						}
					}
					s = append(s, in[i])
				} else {
					s = append(s, in[i])
				}
			}
		}
	}
	for len(s) > 0 {
		top := s[len(s)-1]
		s = s[:len(s)-1]
		post = append(post, top)
	}

	var stack []int
	for i := 0; i < len(post); i++ {
		switch post[i].(type) {
		case int:
			{
				stack = append(stack, post[i].(int))
			}
		case byte:
			{
				n1 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				n2 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if post[i].(byte) == '+' {
					stack = append(stack, n1+n2)
				} else if post[i].(byte) == '-' {
					stack = append(stack, n2-n1)
				} else if post[i].(byte) == '*' {
					stack = append(stack, n1*n2)
				} else {
					stack = append(stack, n2/n1)
				}
			}
		}
	}

	return stack[0]
}
