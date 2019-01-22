package _62

import "testing"

type mn struct {
	M int
	N int
}

func TestUniquePaths(t *testing.T) {
	mns := []mn{
		{M: 3, N: 2},
		{M: 7, N: 3},
	}
	expects := []int{
		3,
		28,
	}
	for index, mn := range mns {
		actual := uniquePaths(mn.M, mn.N)
		expect := expects[index]
		if expect != actual {
			t.Fatalf("tests %v ,expect %v, actual %v", mn, expect, actual)
		}
	}
}
