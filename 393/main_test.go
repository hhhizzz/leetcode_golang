package _393

import (
	"fmt"
	"testing"
)

func TestUtf(t *testing.T) {
	var result bool
	result = validUtf8([]int{197, 130, 1})
	fmt.Println(result)
	result = validUtf8([]int{235, 140, 4})
	fmt.Println(result)
	result = validUtf8([]int{237})
	fmt.Println(result)
	result = validUtf8([]int{240, 162, 138, 147})
	fmt.Println(result)
}
