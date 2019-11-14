package _621

import (
    "fmt"
    "testing"
)

func TestLeastInterval(t *testing.T) {
    var result int

    result = leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2)
    fmt.Println(result)

    result = leastInterval([]byte{'A', 'B', 'C', 'D', 'E', 'A', 'B', 'C', 'D', 'E'}, 4)
    fmt.Println(result)
}
