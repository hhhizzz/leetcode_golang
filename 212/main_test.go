package _212

import (
    "fmt"
    "testing"
)

func TestConstruct(t *testing.T) {
    trie := construct(&[]string{"hello", "world"})
    println(trie)
}

func TestExist(t *testing.T) {
    board := [][]byte{
        {'o', 'a', 'a', 'n'},
        {'e', 't', 'a', 'e'},
        {'i', 'h', 'k', 'r'},
        {'i', 'f', 'l', 'v'}}
    expect := []string{"oath", "pea", "eat", "rain"}
    fmt.Println(findWords(board, expect))
}
