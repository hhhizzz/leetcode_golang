package _253

import (
	"fmt"
	"testing"
)

func TestMinMeetingRooms(t *testing.T) {
	var result int
	result = minMeetingRooms([][]int{{7, 10}, {2, 4}})
	fmt.Println(result)
}
