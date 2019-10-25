package _443

import (
    "fmt"
    "testing"
)

func TestCompress(t *testing.T) {
    input := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'c', 'c', 'c'}
    length := compress(input)
    fmt.Println(length)
    for i := 0; i < length; i++ {
        fmt.Printf("%c, ", input[i])
    }

}
