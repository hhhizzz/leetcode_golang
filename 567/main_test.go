package _567

import (
	"fmt"
	"testing"
)

func TestCheckInclusion(t *testing.T) {
	var result bool
	result = checkInclusion("ab", "eidbaooo")
	fmt.Println(result)

	result = checkInclusion("ab", "eidboaoo")
	fmt.Println(result)

	result = checkInclusion("pqzhi", "ghrqpihzybre")
	fmt.Println(result)

	result = checkInclusion("abbc", "bbabbaa")
	fmt.Println(result)
}
