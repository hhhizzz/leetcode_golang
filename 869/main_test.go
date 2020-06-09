package _869

import (
	"fmt"
	"testing"
)

func TestToByte(t *testing.T) {
	fmt.Println(toByte(1))
	fmt.Println(toByte(12))
	fmt.Println(toByte(10))
	fmt.Println(toByte(100))
	fmt.Println(toByte(1000))
}

func TestQSort(t *testing.T) {
	input := []byte{9, 8, 7, 6, 5, 4, 3, 2, 1}
	qSort(input)
	fmt.Println(input)
}

func TestReorderedPowerOf2(t *testing.T) {
	inputs := []int{
		1,
		10,
		16,
		24,
		46,
		218,
	}
	expects := []bool{
		true,
		false,
		true,
		false,
		true,
		true,
	}
	for i, expect := range expects {
		actual := reorderedPowerOf2(inputs[i])
		if actual != expect {
			t.Errorf("input %v,expected %v, actual %v", inputs[i], expect, actual)
		}
	}
}
