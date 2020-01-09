package _65

import "regexp"

func isNumber(s string) bool {
    pattern, _ := regexp.Compile(`^(\s)*([+\-])?((\d)+(\.\d*)?|(\.\d+))(e([+\-])?\d+)?(\s)*$`)
    return pattern.Match([]byte(s))
}
