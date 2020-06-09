package _150

import (
	"fmt"
	"testing"
)

func TestEvalRPN(t *testing.T) {
	rpn := evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})
	fmt.Println(rpn) //22
}
