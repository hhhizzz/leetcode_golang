package _295

import (
	"fmt"
	"testing"
)

func TestFindMedian(t *testing.T) {
	obj := Constructor()
	obj.AddNum(1)
	fmt.Println(obj.FindMedian())
	obj.AddNum(2)
	fmt.Println(obj.FindMedian())
	obj.AddNum(3)
	fmt.Println(obj.FindMedian())
	obj.AddNum(4)
	fmt.Println(obj.FindMedian())
	obj.AddNum(5)
	fmt.Println(obj.FindMedian())
}
