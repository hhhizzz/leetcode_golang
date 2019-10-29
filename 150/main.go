package _150

import "strconv"

type stack []int

func (s *stack) push(item int) {
    *s = append(*s, item)
}

func (s *stack) pop() int {
    result := (*s)[len(*s)-1]
    *s = (*s)[:len(*s)-1]
    return result
}

func evalRPN(tokens []string) int {
    s := stack{}
    for _, token := range tokens {
        v, err := strconv.Atoi(token)
        if err == nil {
            s.push(v)
        } else {
            b := s.pop()
            a := s.pop()
            var c int
            if token == "+" {
                c = a + b
            } else if token == "-" {
                c = a - b
            } else if token == "*" {
                c = a * b
            } else {
                c = a / b
            }
            s.push(c)
        }
    }
    return s.pop()
}
