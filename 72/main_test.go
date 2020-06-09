package _72

import (
	"fmt"
	"testing"
)

func TestMinDistance(t *testing.T) {
	fmt.Println(minDistance("intention", "execution"))
	fmt.Println(minDistance("", ""))
	fmt.Println(minDistance("", "123"))
	fmt.Println(minDistance("pneumonoultramicroscopicsilicovolcanoconiosis", "ultramicroscopically"))
	fmt.Println(minDistance("ii", "iiii"))
}
