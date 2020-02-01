package _402

import (
    "fmt"
    "testing"
)

func TestRemoveKdigits(t *testing.T) {
    var result string

    result = removeKdigits("1432219", 3)
    fmt.Println(result)

    result = removeKdigits("1230", 3)
    fmt.Println(result)

    result = removeKdigits("12", 1)
    fmt.Println(result)

    result = removeKdigits("10200", 1)
    fmt.Println(result)

    result = removeKdigits("10200", 3)
    fmt.Println(result)
}
