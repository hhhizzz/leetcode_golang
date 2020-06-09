package _52

import (
	"fmt"
	"testing"
	"time"
)

func TestTotalNQueens(t *testing.T) {
	for i := 4; i < 20; i++ {
		t1 := time.Now()
		j := totalNQueens(i)
		t2 := time.Now()
		fmt.Printf("the number of %d is %d\n", i, j)
		fmt.Printf("costs time %v\n\n", t2.Sub(t1))
	}

}
