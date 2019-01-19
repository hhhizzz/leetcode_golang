package _56

import "testing"

func TestMerge(t *testing.T) {
	//[[1,3],[2,6],[8,10],[15,18]]
	intervals := []Interval{
		{1, 3},
		{2, 6},
		//{8, 10},
		//{15, 18},
	}
	expect := []Interval{
		{1, 6},
		//{8, 10},
		//{15, 18},
	}
	actual := merge(intervals)
	if len(expect) != len(actual) {
		t.Fatal("长度错误")
	}
	for i, a := range actual {
		if expect[i] != a {
			t.Fatalf("expect: %v,actual: %v", expect, actual)
		}
	}
}
