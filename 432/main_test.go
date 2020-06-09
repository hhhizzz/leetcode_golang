package _432

import (
	"fmt"
	"testing"
)

func TestAllOne(t *testing.T) {
	all := Constructor()
	all.Inc("a")
	all.Inc("a")
	all.Inc("b")
	all.Inc("b")
	all.Dec("b")
	all.Dec("b")
	all.Inc("b")
	fmt.Println(all.GetMaxKey())
	fmt.Println(all.GetMinKey())
}
