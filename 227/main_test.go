package _227

import (
    "fmt"
    "testing"
)

func TestCal(t *testing.T) {
    var result int

    result = calculate("3+2*2")
    fmt.Println(result) //7

    result = calculate("3+2*2+21*21+213*42")
    fmt.Println(result) //9394

    result = calculate(" 3+5 / 2 ")
    fmt.Println(result) //5
}
