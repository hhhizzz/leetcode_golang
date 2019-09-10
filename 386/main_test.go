package _386

import (
    "fmt"
    "testing"
)

func TestLexicalOrder(t *testing.T) {
    order := lexicalOrder(13)
    fmt.Println(order)
}
