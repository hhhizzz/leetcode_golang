package _212

import (
    "fmt"
    "testing"
)

func TestExist(t *testing.T) {
    board := [][]byte{{'s', 'e', 'e', 'n', 'e', 'w'}, {'t', 'm', 'r', 'i', 'v', 'a'}, {'o', 'b', 's', 'i', 'b', 'd'}, {'w', 'm', 'y', 's', 'e', 'n'}, {'l', 't', 's', 'n', 's', 'a'}, {'i', 'e', 'z', 'l', 'g', 'n'}}
    expect := []string{"anda", "anes", "anesis", "avener", "avine", "bena", "bend", "benda", "besa", "besan", "bowl", "daven", "embow", "inerm", "irene", "myst", "nane", "nanes", "neem", "reem", "reest", "renew", "rine", "riva", "rive", "riven", "sand", "sane", "sang", "seen", "seer", "send", "sise", "stob", "stow", "teil", "vine", "viner", "vire", "wadna", "wave", "wene", "wots"}
    fmt.Println(findWords(board, expect))
}
