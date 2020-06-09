package _892

import (
	"fmt"
	"testing"
)

func TestSurfaceArea(t *testing.T) {
	area := surfaceArea([][]int{{1, 0}, {0, 2}})
	fmt.Println(area)
}
